package queue

import "fmt"

type QueueNode struct {
	Val  interface{}
	Next *QueueNode
}

type QueueLink struct {
	Front  *QueueNode
	Rear   *QueueNode
	Length int
}

func InitQueueLink() *QueueLink {
	return &QueueLink{
		Front:  nil,
		Rear:   nil,
		Length: 0,
	}
}

func (l *QueueLink) Empty() bool {
	return l.Length == 0
}

func (l *QueueLink) EnQueue(val interface{}) {
	node := &QueueNode{
		Val:  val,
		Next: nil,
	}
	if l.Empty() {
		l.Front = node
		l.Rear = node
		l.Length = 1
	} else {
		l.Rear.Next = node
		l.Rear = node
		l.Length++
	}
}

func (l *QueueLink) DeQueue() *QueueNode {
	if l.Empty() {
		return nil
	}
	val := l.Front
	if l.Front == l.Rear { //只有一个元素
		l.Front = nil
		l.Rear = nil
		l.Length = 0
	} else {
		pre := l.Front
		l.Front = pre.Next
		l.Length--
		pre = nil
	}
	return val
}

func (l *QueueLink) List() {
	if l.Empty() {
		return
	}
	cur := l.Front
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
