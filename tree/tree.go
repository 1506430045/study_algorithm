package tree

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"study_algorithm/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) []int {
	res := make([]int, 0)
	DoPreOrder(root, &res)
	return res
}

func DoPreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	DoPreOrder(root.Left, res)
	DoPreOrder(root.Right, res)
	*res = append(*res, root.Val)
}

func PreOrder2(root *TreeNode) {
	if root == nil {
		return
	}
	s := stack.InitStack()
	s.Push(root)

	for !s.Empty() {
		node := s.Pop().(*TreeNode)
		fmt.Print(node.Val, "->")
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

func InOrder2(root *TreeNode) {
	if root == nil {
		return
	}
	s := stack.InitStack()
	cur := root
	for !s.Empty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			cur = s.Pop().(*TreeNode)
			fmt.Print(cur.Val, "->")
			cur = cur.Right
		}
	}
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s1 := stack.InitStack()
	s2 := stack.InitStack()
	s1.Push(root)
	for !s1.Empty() {
		node := s1.Pop().(*TreeNode)
		s2.Push(node)
		if node.Left != nil {
			s1.Push(node.Left)
		}
		if node.Right != nil {
			s1.Push(node.Right)
		}
	}

	for !s2.Empty() {
		fmt.Print(s2.Pop().(*TreeNode).Val, "->")
	}
}

func LevelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	q := InitQueueLink()
	q.EnQueue(root)
	for !q.Empty() {
		node := q.DeQueue().Val.(*TreeNode)
		fmt.Print(node.Val, "->")
		if node.Left != nil {
			q.EnQueue(node.Left)
		}
		if node.Right != nil {
			q.EnQueue(node.Right)
		}
	}
}

func levelOrder(root *TreeNode) [][]int {
	var result = make([][]int, 0)
	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	q := InitQueueLink()
	q.EnQueue(root)
	var level = 0
	var length = 0
	for !q.Empty() {
		length = q.Length
		level++
		for length > 0 {
			length--
			node := q.DeQueue()
			if node == nil {
				break
			}
			queueNode := node.Val.(*TreeNode)
			if queueNode.Left != nil {
				if len(result) <= level {
					result = append(result, []int{})
				}
				result[level] = append(result[level], queueNode.Left.Val)
				q.EnQueue(queueNode.Left)
			}
			if queueNode.Right != nil {
				if len(result) <= level {
					result = append(result, []int{})
				}
				result[level] = append(result[level], queueNode.Right.Val)
				q.EnQueue(queueNode.Right)
			}
		}
	}
	return result
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

func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := MaxDepth2(root.Left) + 1
	rightMax := MaxDepth2(root.Right) + 1
	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}

func MaxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	q := InitQueueLink()
	q.EnQueue(root)
	for !q.Empty() {
		length := q.Length
		for length > 0 {
			length--
			node := q.DeQueue()
			if node == nil {
				break
			}
			queueNode := node.Val.(*TreeNode)
			if queueNode.Left != nil {
				q.EnQueue(queueNode.Left)
			}
			if queueNode.Right != nil {
				q.EnQueue(queueNode.Right)
			}
		}
		depth++
	}
	return depth
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return nil
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i <= length-3; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := twoSum(nums[i+1:length], 0-nums[i])
		if len(r) == 0 {
			continue
		}
		for _, rr := range r {
			if len(rr) == 2 {
				res = append(res, []int{nums[i], rr[0], rr[1]})
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	left := 0
	right := len(nums) - 1

	m := make(map[int]bool)
	res := make([][]int, 0)
	sum := 0
	for left < right {
		sum = nums[left] + nums[right]
		if sum == target && !m[nums[left]] {
			res = append(res, []int{nums[left], nums[right]})
			m[nums[left]] = true
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return res
}

func isUnique(astr string) bool {
	m := make(map[uint8]bool)

	for i := 0; i < len(astr); i++ {
		if _, ok := m[astr[i]]; ok {
			return false
		}
		m[astr[i]] = true
	}
	return true
}

func CheckPermutation(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return false
	}
	m := make(map[uint8]uint8)
	for i := 0; i < l1; i++ {
		if _, ok := m[s1[i]]; ok {
			m[s1[i]]++
		} else {
			m[s1[i]] = 1
		}
	}
	for i := 0; i < l2; i++ {
		if n, ok := m[s2[i]]; ok {
			m[s2[i]]--
			if n <= 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func replaceSpaces(S string, length int) string {
	l := len(S)

	var i int
	var build strings.Builder
	for i < l && i < length {
		if S[i] == 32 {
			build.WriteString("%20")
		} else {
			build.WriteByte(S[i])
		}
		i++
	}
	return build.String()
}

func canPermutePalindrome(s string) bool {
	m := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			delete(m, s[i])
		} else {
			m[s[i]] = 1
		}
	}
	return len(m) <= 1
}

//aabcccccaaa
func compressString(S string) string {
	length := len(S)
	if length <= 2 {
		return S
	}
	build := strings.Builder{}

	f := S[0]
	count := 1
	for i := 1; i < length; i++ {
		if S[i] != f {
			build.WriteByte(f)
			build.WriteString(strconv.Itoa(count))

			f = S[i]
			count = 1
			continue
		}
		count++
	}
	build.WriteByte(f)
	build.WriteString(strconv.Itoa(count))
	str := build.String()
	if len(str) >= len(S) {
		return S
	}
	return str
}

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1 += s1
	return strings.Index(s1, s2) != -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		find := false
		n := -1
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				find = true
				continue
			}
			if !find {
				continue
			}
			if nums2[j] > nums1[i] {
				n = nums2[j]
				break
			}
		}
		res = append(res, n)
	}
	return res
}

//莫里斯先序遍历

//1. 如果cur无左孩子，cur向右移动（cur=cur.right）
//2. 如果cur有左孩子，找到cur左子树上最右的节点，记为mostright
//2.1 如果mostright的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
//2.2 如果mostright的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）
func morrisPre(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var mostRight *TreeNode
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				fmt.Print(cur.Val, " ")
				cur = cur.Left
				continue
			} else {
				mostRight = nil
			}
		} else {
			fmt.Print(cur.Val, " ")
		}
		cur = cur.Right
	}
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := INT_MIN
	oneSideMax(root, &res)
	return res
}

func oneSideMax(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	var left, right int
	left = max(0, oneSideMax(root.Left, res))
	right = max(0, oneSideMax(root.Right, res))
	*res = max(*res, left+right+root.Val)
	fmt.Println(*res)

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func ConvertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, rear *TreeNode
	s := stack.InitStack()
	cur := root
	for !s.Empty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			node := s.Pop().(*TreeNode)
			resNode := &TreeNode{
				Val:   node.Val,
				Left:  nil,
				Right: nil,
			}
			if head == nil {
				head = resNode
				rear = head
			} else {
				rear.Right = resNode
				rear = rear.Right
			}
			cur = node.Right
		}
	}
	return head
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		}
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur
}

func isBalanced(root *TreeNode) bool {
	flag := true
	dfs(root, &flag)
	return flag
}

func dfs(node *TreeNode, flag *bool) int {
	if node == nil || !*flag {
		return 0
	}
	leftDepth := dfs(node.Left, flag) + 1
	rightDepth := dfs(node.Right, flag) + 1
	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		*flag = false
	}
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

func InvertTree(root *TreeNode) *TreeNode {
	doInvertTree(root)
	return root
}

func doInvertTree(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	doInvertTree(root.Left)
	doInvertTree(root.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
	doMirrorTree(root)
	return root
}

func doMirrorTree(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	doMirrorTree(root.Left)
	doMirrorTree(root.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return checkSymmetric(l.Left, r.Right) && checkSymmetric(l.Right, r.Left)
}

func levelOrder2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := InitQueueLink()
	q.EnQueue(root)
	for !q.Empty() {
		node := q.DeQueue().Val.(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			q.EnQueue(node.Left)
		}
		if node.Right != nil {
			q.EnQueue(node.Right)
		}
	}
	return res
}
