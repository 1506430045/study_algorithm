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
