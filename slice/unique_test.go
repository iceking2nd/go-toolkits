package slice

import (
	"fmt"
	"testing"
)

func TestSliceUnique(t *testing.T) {
	s1 := []string{"a", "b", "c", "g", "g", "a", "c"}
	s2 := []int{1, 3, 9, 2, 5, 7, 3, 9, 4, 4, 1}
	s3 := []bool{true, false, true, false, false, false, true, true}

	fmt.Printf("s1 origin: %v\n", s1)
	fmt.Printf("s1 uniqued: %v\n", Unique(s1))

	fmt.Printf("s2 origin: %v\n", s2)
	fmt.Printf("s2 uniqued: %v\n", Unique(s2))

	fmt.Printf("s3 origin: %v\n", s3)
	fmt.Printf("s3 uniqued: %v\n", Unique(s3))
}
