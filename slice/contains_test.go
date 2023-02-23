package slice

import (
	"fmt"
	"testing"
)

func TestSliceContains(t *testing.T) {
	a := "asdhkjasdhkja"
	b := "aaa"
	fmt.Println(Contains(a, b))
	c := []string{"aa", "bb", "cc"}
	d := "aaa"
	fmt.Println(Contains(c, d))
	e := []string{"aa", "bb", "cc"}
	f := []string{"aa"}
	fmt.Println(Contains(e, f))
	g := []string{"aa", "bb", "cc"}
	h := []uint64{1}
	fmt.Println(Contains(g, h))
	i := []interface{}{1, 2, 3}
	j := uint64(3)
	fmt.Println(Contains(i, j))
}
