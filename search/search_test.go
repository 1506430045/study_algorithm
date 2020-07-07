package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1,0,3,5,9,12}
	target := 5
	i := search(nums, target)
	fmt.Println(i)
}

func TestMySqrt(t *testing.T)  {
	x := mySqrt(0)
	fmt.Println(x)
}
