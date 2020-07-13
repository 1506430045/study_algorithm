package list

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	a := &ListNode{
		Val:  1,
		Next: nil,
	}
	b := &ListNode{
		Val:  2,
		Next: nil,
	}
	c := &ListNode{
		Val:  4,
		Next: nil,
	}

	d := &ListNode{
		Val:  1,
		Next: nil,
	}
	e := &ListNode{
		Val:  3,
		Next: nil,
	}
	f := &ListNode{
		Val:  4,
		Next: nil,
	}

	a.Next = b
	b.Next = c

	d.Next = e
	e.Next = f

	xx := mergeTwoLists(a, nil)

	cur := xx
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
