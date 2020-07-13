package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, rear *ListNode
	var sum int
	var upFlag = false
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if upFlag {
			sum += 1
		}
		if sum > 9 {
			upFlag = true
			sum = sum - 10
		} else {
			upFlag = false
		}
		if head == nil || rear == nil {
			head = &ListNode{
				Val:  sum,
				Next: nil,
			}
			rear = head
		} else {
			rear.Next = &ListNode{
				Val:  sum,
				Next: nil,
			}
			rear = rear.Next
		}
	}
	if upFlag && rear != nil {
		rear.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return head
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := make(map[int]bool)
	cur := head
	pre := head
	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			pre.Next = cur.Next
			pre = cur.Next
			cur = cur.Next
		} else {
			pre = cur
		}
		if cur.Next != nil {
			cur = cur.Next
		}
		m[cur.Val] = true
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}
	cur := head.Next.Next

	head.Next.Next = head
	head.Next = cur
	return swapPairs(cur)
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head

	for fast != nil && slow != nil {
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA := headA
	curB := headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	return false
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	length := len(res)
	for i := 0; i < length/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return res
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	cur := head
	prev := head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		}
		prev = cur
		cur = cur.Next
	}
	return head
}

func reverseList(head *ListNode) *ListNode {
	var prev, temp *ListNode
	var cur = head

	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, rear, node *ListNode
	var v int
	for l1 != nil || l2 != nil {

		if l1 != nil && l2 != nil { //l1、l2都不为nil 选择小的
			if l1.Val <= l2.Val {
				v = l1.Val
				l1 = l1.Next
			} else {
				v = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil { //l1不为nil 则l2为nil
			v = l1.Val
			l1 = l1.Next
		} else if l2 != nil { //l2不为nil 则l1为nil
			v = l2.Val
			l2 = l2.Next
		}
		node = &ListNode{
			Val:  v,
			Next: nil,
		}
		if head == nil || rear == nil {
			head = node
			rear = node
		} else {
			rear.Next = node
			rear = rear.Next
		}
	}
	return head
}