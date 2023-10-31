package addtwonumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var prev *ListNode
	var carry, value int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1, l2 = l2, l1
		}
		if l2 == nil {
			value = l1.Val
		} else {
			value = l1.Val + l2.Val
		}
		value += carry
		if value >= 10 {
			carry = 1
			value -= 10
		} else {
			carry = 0
		}
		if prev == nil {
			prev = &ListNode{Val: value}
			head = prev
		} else {
			prev.Next = &ListNode{Val: value}
			prev = prev.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		prev.Next = &ListNode{Val: carry}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
