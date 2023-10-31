package reverse_nodes_in_k_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	h := createLinkedList([]int{1, 2, 3, 4})
	newHead := reverseKGroup(h, 1)
	assert.Equal(t, []int{1, 2, 3, 4}, listToArray(newHead))

	h = createLinkedList([]int{1, 2, 3, 4})
	newHead = reverseKGroup(h, 2)
	assert.Equal(t, []int{2, 1, 4, 3}, listToArray(newHead))

	h = createLinkedList([]int{1, 2, 3, 4})
	newHead = reverseKGroup(h, 3)
	assert.Equal(t, []int{3, 2, 1, 4}, listToArray(newHead))

	h = createLinkedList([]int{1, 2, 3, 4})
	newHead = reverseKGroup(h, 4)
	assert.Equal(t, []int{4, 3, 2, 1}, listToArray(newHead))
}

func createLinkedList(values []int) *ListNode {
	n := len(values)
	head := &ListNode{
		Val:  values[n-1],
		Next: nil,
	}
	for i := n - 2; i >= 0; i-- {
		newNode := &ListNode{
			Val:  values[i],
			Next: head,
		}
		head = newNode
	}
	return head
}

func listToArray(h *ListNode) []int {
	var results []int
	for h != nil {
		results = append(results, h.Val)
		h = h.Next
	}
	return results
}
