package heap

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}
