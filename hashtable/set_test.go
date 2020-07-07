package hashtable

import (
	"fmt"
	"testing"
)

func TestMyHashSet_Add(t *testing.T) {
	h := Constructor()
	h.Add(1)
	h.Add(3)
	h.Add(10001)
	fmt.Println(h.Contains(1))
	fmt.Println(h.Contains(2))
	fmt.Println(h.Contains(3))
	h.Remove(3)
	h.Add(10003)
	fmt.Println(h.Contains(3))
	fmt.Println(h.Contains(10001))

	fmt.Print(h.m)
}
