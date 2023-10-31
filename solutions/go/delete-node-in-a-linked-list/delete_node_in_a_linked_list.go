package delete_node_in_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for curr := node; curr.Next != nil; curr = curr.Next {
		next := curr.Next
		curr.Val = next.Val
		if next.Next == nil {
			curr.Next = nil
			return
		}
	}
}
