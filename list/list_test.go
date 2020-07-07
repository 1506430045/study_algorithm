package list

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	a := &ListNode{
		Val:  9,
		Next: nil,
	}
	//b := &ListNode{
	//	Val:  8,
	//	Next: nil,
	//}
	//c := &ListNode{
	//	Val:  8,
	//	Next: nil,
	//}
	//d := &ListNode{
	//	Val:  8,
	//	Next: nil,
	//}
	//e := &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}

	a1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	//b1 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//c1 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//d1 := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//e1 := &ListNode{
	//	Val:  8,
	//	Next: nil,
	//}

	//a.Next = b
	//b.Next = c
	//c.Next = d
	//d.Next = e

	//a1.Next = b1
	//b1.Next = c1
	//c1.Next = d1
	//d1.Next = e1

	aa := addTwoNumbers(a, a1)

	cur := aa
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}

}
