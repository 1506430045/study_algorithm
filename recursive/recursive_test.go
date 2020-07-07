package recursive

import (
	"fmt"
	"testing"
)

func TestPrintReverse(t *testing.T) {
	printReverse("abc")

	fmt.Println()
}

func TestReverseString(t *testing.T) {
	s := []byte{65, 66, 67}
	reverseString(s)

	fmt.Println(s)
}
