package slice

import (
	"fmt"
	"testing"
)

func TestSliceDiff(t *testing.T) {
	s1 := []int{1, 2, 4, 65, 7, 34}
	s2 := []int{2, 4, 8, 9, 15, 23}
	fmt.Println(Diff(s1, s2))
	s3 := []int{1, 2, 4, 65, 7, 34}
	s4 := "adsadasdasd"
	fmt.Println(Diff(s3, s4))
	s5 := []int{1, 2, 4, 65, 7, 34}
	s6 := []uint{2, 4, 8, 9, 15, 23}
	fmt.Println(Diff(s5, s6))
	s7 := []string{"aa", "bb", "cc", "ee"}
	s8 := []string{"aa", "cc", "dd"}
	fmt.Println(Diff(s7, s8))
	s9 := []int{1, 2, 4, 65, 7, 34}
	s10 := []int{1, 2, 4, 65, 7, 34}
	fmt.Println(Diff(s9, s10))
}
