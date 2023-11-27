package insertintoabinarysearchtree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}
	if root == nil {
		return newNode
	}
	var parent *TreeNode
	cur := root
	for cur != nil {
		parent = cur
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if val < parent.Val {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
