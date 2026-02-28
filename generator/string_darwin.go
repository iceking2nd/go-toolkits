//go:build darwin

package generator

import (
	"crypto/rand"
	"errors"
	"math/bits"
	"unsafe"
)

const defaultStringCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

////////////////////////////////////////////////////////
// Options
////////////////////////////////////////////////////////

type StringOption func(*StringConfig)

type StringConfig struct {
	charset string
}

// WithStringCharset 设置字符串字符集
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
// Secure String Generator Core
////////////////////////////////////////////////////////

type SecureStringGenerator struct {
	charset       string
	charsetLen    int
	indexBits     int
	indexMask     int
	indicesPerInt int
}

func newSecureStringGenerator(cfg *StringConfig) (*SecureStringGenerator, error) {
	if len(cfg.charset) == 0 {
		return nil, errors.New("charset cannot be empty")
	}

	charsetLen := len(cfg.charset)
	//indexBits := int(math.Ceil(math.Log2(float64(charsetLen))))
	indexBits := bits.Len(uint(charsetLen - 1))
	indexMask := 1<<indexBits - 1
	indicesPerInt := 64 / indexBits

	return &SecureStringGenerator{
		charset:       cfg.charset,
		charsetLen:    charsetLen,
		indexBits:     indexBits,
		indexMask:     indexMask,
		indicesPerInt: indicesPerInt,
	}, nil
}

////////////////////////////////////////////////////////
// SAFE VERSION
////////////////////////////////////////////////////////

// GenerateStringSafe 生成 crypto 安全字符串（安全转换）
func GenerateStringSafe(n int, opts ...StringOption) (string, error) {
	if n <= 0 {
		return "", nil
	}

	cfg := newStringConfig(opts...)
	gen, err := newSecureStringGenerator(cfg)
	if err != nil {
		return "", err
	}

	b := make([]byte, n)

	randomBuf := make([]byte, 8)
	var cache uint64
	remain := 0

	for i := 0; i < n; {
		if remain == 0 {
			if _, err := rand.Read(randomBuf); err != nil {
				return "", err
			}
			cache = uint64(randomBuf[0]) |
				uint64(randomBuf[1])<<8 |
				uint64(randomBuf[2])<<16 |
				uint64(randomBuf[3])<<24 |
				uint64(randomBuf[4])<<32 |
				uint64(randomBuf[5])<<40 |
				uint64(randomBuf[6])<<48 |
				uint64(randomBuf[7])<<56

			remain = gen.indicesPerInt
		}

		idx := int(cache & uint64(gen.indexMask))
		if idx < gen.charsetLen {
			b[i] = gen.charset[idx]
			i++
		}

		cache >>= gen.indexBits
		remain--
	}

	return string(b), nil
}

////////////////////////////////////////////////////////
// UNSAFE VERSION
////////////////////////////////////////////////////////

// GenerateStringUnsafe 生成 crypto 安全字符串（零拷贝版本）
func GenerateStringUnsafe(n int, opts ...StringOption) (string, error) {
	if n <= 0 {
		return "", nil
	}

	cfg := newStringConfig(opts...)
	gen, err := newSecureStringGenerator(cfg)
	if err != nil {
		return "", err
	}

	b := make([]byte, n)

	randomBuf := make([]byte, 8)
	var cache uint64
	remain := 0

	for i := 0; i < n; {
		if remain == 0 {
			if _, err := rand.Read(randomBuf); err != nil {
				return "", err
			}
			cache = uint64(randomBuf[0]) |
				uint64(randomBuf[1])<<8 |
				uint64(randomBuf[2])<<16 |
				uint64(randomBuf[3])<<24 |
				uint64(randomBuf[4])<<32 |
				uint64(randomBuf[5])<<40 |
				uint64(randomBuf[6])<<48 |
				uint64(randomBuf[7])<<56

			remain = gen.indicesPerInt
		}

		idx := int(cache & uint64(gen.indexMask))
		if idx < gen.charsetLen {
			b[i] = gen.charset[idx]
			i++
		}

		cache >>= gen.indexBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b)), nil
}
