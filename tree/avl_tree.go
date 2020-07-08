package tree

//https://studygolang.com/articles/16735

type AvlTreeNode struct {
	Val   int
	High  int
	Left  *AvlTreeNode
	Right *AvlTreeNode
}

func (this *AvlTreeNode) Insert(v int) *AvlTreeNode {
	if this == nil {
		return &AvlTreeNode{v, 0, nil, nil}
	}
	if v < this.Val {
		this.Left = this.Left.Insert(v)
		this.High = getMax(this.Left.getNodeHigh(), this.Right.getNodeHigh()) + 1
	} else {

	}
	return this
}

func (this *AvlTreeNode) getNodeHigh() int {
	if this == nil {
		return -1
	}
	return this.High
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
