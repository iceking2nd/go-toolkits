package generator

import (
	"strings"
	"sync"
	"testing"
)

func TestGenerateStringSafe_Default(t *testing.T) {
	s, err := GenerateStringSafe(32)
	if err != nil {
		t.Fatal(err)
	}

	if len(s) != 32 {
		t.Fatalf("expected length 32, got %d", len(s))
	}
}

func TestGenerateStringUnsafe_Default(t *testing.T) {
	s, err := GenerateStringUnsafe(32)
	if err != nil {
		t.Fatal(err)
	}

	if len(s) != 32 {
		t.Fatalf("expected length 32, got %d", len(s))
	}
}

func TestGenerateString_CustomCharset(t *testing.T) {
	charset := "0123456789"

	s, err := GenerateStringSafe(100, WithStringCharset(charset))
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range s {
		if !strings.ContainsRune(charset, c) {
			t.Fatalf("invalid character found: %c", c)
		}
	}
}

func TestGenerateString_EmptyLength(t *testing.T) {
	s, err := GenerateStringSafe(0)
	if err != nil {
		t.Fatal(err)
	}

	if s != "" {
		t.Fatalf("expected empty string, got %s", s)
	}
}

func TestGenerateString_InvalidCharset(t *testing.T) {
	_, err := GenerateStringSafe(10, WithStringCharset(""))
	if err != nil {
		t.Fatal("empty charset should fallback to default")
	}
}

func TestGenerateString_Concurrency(t *testing.T) {
	wg := sync.WaitGroup{}
	workers := 100

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s, err := GenerateStringSafe(64)
			if err != nil {
				t.Error(err)
			}
			if len(s) != 64 {
				t.Errorf("wrong length: %d", len(s))
			}
		}()
	}

	wg.Wait()
}

func TestGenerateString_UniqueProbability(t *testing.T) {
	set := make(map[string]struct{})
	iterations := 10000

	for i := 0; i < iterations; i++ {
		s, err := GenerateStringSafe(16)
		if err != nil {
			t.Fatal(err)
		}
		set[s] = struct{}{}
	}

	if len(set) < iterations-5 {
		t.Fatalf("too many duplicates: %d/%d", len(set), iterations)
	}
}

func BenchmarkGenerateStringSafe_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateStringSafe(16)
	}
}

func BenchmarkGenerateStringUnsafe_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateStringUnsafe(16)
	}
}

func BenchmarkGenerateStringSafe_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateStringSafe(64)
	}
}

func BenchmarkGenerateStringUnsafe_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateStringUnsafe(64)
	}
}

func BenchmarkGenerateStringSafe_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateStringSafe(256)
	}
}

func BenchmarkGenerateStringUnsafe_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateStringUnsafe(256)
	}
}

func BenchmarkGenerateStringSafe_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = GenerateStringSafe(32)
		}
	})
}
