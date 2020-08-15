package generator

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(20))
	fmt.Println(RandStringRunes(20))
	fmt.Println(RandStringBytes(20))
	fmt.Println(RandStringBytesRmndr(20))
	fmt.Println(RandStringBytesMask(20))
	fmt.Println(RandStringBytesMaskImpr(20))
	fmt.Println(RandStringBytesMaskImprSrc(20))
	fmt.Println(RandStringBytesMaskImprSrcSB(20))
	fmt.Println(RandStringBytesMaskImprSrcUnsafe(20))
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(20)
	}
}

func BenchmarkRandStringRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes(20)
	}
}

func BenchmarkRandStringBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytes(20)
	}
}

func BenchmarkRandStringBytesRmndr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesRmndr(20)
	}
}

func BenchmarkRandStringBytesMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMask(20)
	}
}

func BenchmarkRandStringBytesMaskImpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpr(20)
	}
}

func BenchmarkRandStringBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrc(20)
	}
}

func BenchmarkRandStringBytesMaskImprSrcSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcSB(20)
	}
}

func BenchmarkRandStringBytesMaskImprSrcUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrcUnsafe(20)
	}
}
