package other

type Pos struct {
	X int
	Y int
}

func numIslands(grid [][]byte) int {
	var i, j, nX, nY, numLands int
	var q *QueueLink

	nX = len(grid)
	nY = len(grid[0])

	q = InitQueueLink()

	for i = 0; i < nX; i++ {
		for j = 0; j < nY; j++ {
			if grid[i][j] == 1 { //陆地
				numLands++
				grid[i][j] = 0
				q.EnQueue(&Pos{i, j})

				for !q.Empty() {
					node := q.DeQueue().Val.(*Pos)
					if node.X-1 >= 0 && grid[node.X-1][node.Y] == 1 { //左边也为陆地
						grid[node.X-1][node.Y] = 0
						q.EnQueue(&Pos{node.X - 1, node.Y})
					}
					if node.X+1 < nX && grid[node.X+1][node.Y] == 1 { //右边边也为陆地
						grid[node.X+1][node.Y] = 0
						q.EnQueue(&Pos{node.X + 1, node.Y})
					}
					if node.Y-1 >= 0 && grid[node.X][node.Y-1] == 1 { //上边也为陆地
						grid[node.X][node.Y-1] = 0
						q.EnQueue(&Pos{node.X, node.Y - 1})
					}
					if node.Y+1 < nY && grid[node.X][node.Y+1] == 1 { //下边也为陆地
						grid[node.X][node.Y+1] = 0
						q.EnQueue(&Pos{node.X, node.Y + 1})
					}
				}
			}
		}
	}
	return 1
}

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
