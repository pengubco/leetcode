package validate_binary_search_tree

// https://leetcode.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	valid, _, _ := dfs(root)
	return valid
}

func dfs(cur *TreeNode) (bool, int, int) {
	if cur.Left == nil && cur.Right == nil {
		return true, cur.Val, cur.Val
	}
	curMinVal, curMaxVal := cur.Val, cur.Val
	if cur.Left != nil {
		valid, minVal, maxVal := dfs(cur.Left)
		if !valid || maxVal >= cur.Val {
			return false, 0, 0
		}
		curMinVal = min(curMinVal, minVal)
	}
	if cur.Right != nil {
		valid, minVal, maxVal := dfs(cur.Right)
		if !valid || minVal <= cur.Val {
			return false, 0, 0
		}
		curMaxVal = max(curMaxVal, maxVal)
	}
	return true, curMinVal, curMaxVal
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
