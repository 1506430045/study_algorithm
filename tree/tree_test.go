package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val: 2,
	}
	c := &TreeNode{
		Val: 3,
	}
	d := &TreeNode{
		Val: 4,
	}
	e := &TreeNode{
		Val: 5,
	}
	f := &TreeNode{
		Val: 6,
	}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f

	fmt.Println(maxPathSum(a))
	//fmt.Println(res)
}

func TestMaxDepth(t *testing.T) {
	a := &TreeNode{
		Val: 1,
	}
	//b := &TreeNode{
	//	Val: 2,
	//}
	//c := &TreeNode{
	//	Val: 3,
	//}
	//d := &TreeNode{
	//	Val: 4,
	//}
	//e := &TreeNode{
	//	Val: 5,
	//}
	//f := &TreeNode{
	//	Val: 6,
	//}
	//a.Left = b
	//a.Right = c
	//b.Left = d
	//b.Right = e
	//c.Left = f
	depth := MaxDepth2(a)
	fmt.Println(depth)
}

func TestTwoSum(t *testing.T) {
	//"islander"
	//"slander"
	fmt.Println("xx", nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
