package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// Step 0: Handle special case
	if k == 1 {
		return head
	}

	// Step 1: break into k sublists
	var subLists [][2]*ListNode
	var h *ListNode
	cnt := 1
	h = head
	for curr := head; curr != nil; {
		if cnt == k {
			subLists = append(subLists, [2]*ListNode{h, curr})
			curr = curr.Next
			if curr != nil {
				h = curr
				cnt = 1
			} else {
				cnt = 0
			}
			continue
		}
		curr = curr.Next
		if curr != nil {
			cnt++
		}
	}
	if cnt > 0 {
		subLists = append(subLists, [2]*ListNode{h, nil})
	}

	// Step 2: reverse each sublist
	for _, l := range subLists {
		if l[1] == nil {
			break
		}
		reverse(l[0], l[1])
		l[0].Next = nil
	}

	// Step 3: link reversed sublists.
	newHead := subLists[0][1]
	m := len(subLists)
	for i := 0; i < m-1; i++ {
		if subLists[i+1][1] != nil {
			subLists[i][0].Next = subLists[i+1][1]
		} else {
			subLists[i][0].Next = subLists[i+1][0]
		}
	}

	return newHead
}

// reverse sub link list [h,...,t]. Returns the new head.
func reverse(h, t *ListNode) *ListNode {
	if h == t {
		return h
	}
	tmp := reverse(h.Next, t)
	h.Next.Next = h
	return tmp
}
