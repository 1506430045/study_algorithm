package array

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(Intersect(nums1, nums2))
}

func TestIntersect2(t *testing.T) {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 5, 4, 2, 1}))
}

func TestExchange(t *testing.T) {
	fmt.Println(exchange([]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1}))
}

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2.00000, -100))
}

func TestIsStraight(t *testing.T) {
	fmt.Println(isStraight([]int{0,0,2,2,5}))
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2}, 9))
}
