package populatingnextrightpointersineachnode

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// use two queues. current level and next level
	currQueue, nextQueue := make([]*Node, 0), make([]*Node, 0)
	currQueue = append(currQueue, root)
	for len(currQueue) > 0 {
		for _, currNode := range currQueue {
			if currNode.Left == nil && currNode.Right == nil {
				continue
			}
			nextQueue = append(nextQueue, currNode.Left, currNode.Right)
		}
		n := len(currQueue)
		currQueue[n-1].Next = nil
		for i := n - 2; i >= 0; i-- {
			currQueue[i].Next = currQueue[i+1]
		}
		currQueue, nextQueue = nextQueue, currQueue
		nextQueue = nextQueue[:0]
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
