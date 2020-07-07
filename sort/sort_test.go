package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{2, 1, 8, 4, 3, 7, 6, 9}
	//HeapSort(arr)
	//fmt.Println(arr)


	fmt.Print(getLeastNumbers(arr, 2))
}

func TestQuickSort(t *testing.T) {
	arr := []int{2, 1, 8, 4, 3, 7, 6, 9}
	QuickSort(arr, 0, 7)
	fmt.Println(arr)
}