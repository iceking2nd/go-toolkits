//go:build !darwin

package generator

import (
	"crypto/rand"
	"errors"
	"io"
	"math/bits"
	"sync"
	"unsafe"
)

const defaultStringCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

////////////////////////////////////////////////////////
// Option
////////////////////////////////////////////////////////

type StringOption func(*StringConfig)

type StringConfig struct {
	charset string
}

func WithStringCharset(cs string) StringOption {
	return func(c *StringConfig) {
		if cs != "" {
			c.charset = cs
		}
	}
}

func newStringConfig(opts ...StringOption) *StringConfig {
	cfg := &StringConfig{
		charset: defaultStringCharset,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

////////////////////////////////////////////////////////
// SecureStringGenerator
////////////////////////////////////////////////////////

type secureStringGenerator struct {
	charset       string
	charsetLen    int
	indexBits     int
	indexMask     uint64
	indicesPerInt int
}

func newSecureStringGenerator(cfg *StringConfig) (*secureStringGenerator, error) {
	if len(cfg.charset) == 0 {
		return nil, errors.New("charset cannot be empty")
	}

	charsetLen := len(cfg.charset)

	indexBits := bits.Len(uint(charsetLen - 1))
	indexMask := uint64(1<<indexBits - 1)
	indicesPerInt := 64 / indexBits

	return &secureStringGenerator{
		charset:       cfg.charset,
		charsetLen:    charsetLen,
		indexBits:     indexBits,
		indexMask:     indexMask,
		indicesPerInt: indicesPerInt,
	}, nil
}

////////////////////////////////////////////////////////
// Pool (复用 generator + 随机缓冲区)
////////////////////////////////////////////////////////

type pooledGenerator struct {
	gen   *secureStringGenerator
	randB []byte
}

var generatorPool = sync.Pool{
	New: func() any {
		return &pooledGenerator{
			randB: make([]byte, 4096), // 4KB 批量随机缓存
		}
	},
}

////////////////////////////////////////////////////////
// 内部核心生成逻辑
////////////////////////////////////////////////////////

func generateString(n int, unsafeMode bool, opts ...StringOption) (string, error) {
	if n <= 0 {
		return "", nil
	}

	cfg := newStringConfig(opts...)
	gen, err := newSecureStringGenerator(cfg)
	if err != nil {
		return "", err
	}

	// 从 pool 获取
	pg := generatorPool.Get().(*pooledGenerator)
	pg.gen = gen

	defer func() {
		pg.gen = nil
		generatorPool.Put(pg)
	}()

	b := make([]byte, n)

	// 批量填充随机数
	if _, err := io.ReadFull(rand.Reader, pg.randB); err != nil {
		return "", err
	}

	var cache uint64
	remain := 0
	randIndex := 0
	randLen := len(pg.randB)

	for i := 0; i < n; {
		if remain == 0 {
			if randIndex+8 > randLen {
				// 重新填充 buffer
				if _, err := io.ReadFull(rand.Reader, pg.randB); err != nil {
					return "", err
				}
				randIndex = 0
			}

			// 使用栈 uint64 组合（方案1）
			cache = uint64(pg.randB[randIndex]) |
				uint64(pg.randB[randIndex+1])<<8 |
				uint64(pg.randB[randIndex+2])<<16 |
				uint64(pg.randB[randIndex+3])<<24 |
				uint64(pg.randB[randIndex+4])<<32 |
				uint64(pg.randB[randIndex+5])<<40 |
				uint64(pg.randB[randIndex+6])<<48 |
				uint64(pg.randB[randIndex+7])<<56

			randIndex += 8
			remain = gen.indicesPerInt
		}

		idx := int(cache & gen.indexMask)
		if idx < gen.charsetLen {
			b[i] = gen.charset[idx]
			i++
		}

		cache >>= gen.indexBits
		remain--
	}

	if unsafeMode {
		return *(*string)(unsafe.Pointer(&b)), nil
	}

	return string(b), nil
}

////////////////////////////////////////////////////////
// Public API
////////////////////////////////////////////////////////

// GenerateStringSafe crypto安全 + 内存安全
func GenerateStringSafe(n int, opts ...StringOption) (string, error) {
	return generateString(n, false, opts...)
}

// GenerateStringUnsafe crypto安全 + 零拷贝
func GenerateStringUnsafe(n int, opts ...StringOption) (string, error) {
	return generateString(n, true, opts...)
}
