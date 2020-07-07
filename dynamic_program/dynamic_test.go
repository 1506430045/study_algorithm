package dynamic_program

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	fmt.Println(CoinChange([]int{1, 2, 3}, 6))
}
