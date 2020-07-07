package other

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	cc := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	x := numIslands(cc)

	fmt.Println(x)
}

func TestGenerate(t *testing.T) {
	//res := generate(10)
	//fmt.Println(res)

	fmt.Println(getRow(4))
}

func TestPivotIndex(t *testing.T) {
	fmt.Println(pivotIndex([]int{-1, 1, -2, 2, -4, 2}))
}

func TestSearchInsert(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func TestFindNumberIn2DArray(t *testing.T) {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(arr, 16))
}

func TestReplaceSpace(t *testing.T) {
	fmt.Println(replaceSpace("We are happy."))
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{2, 2, 1, 3, 1}))
}
