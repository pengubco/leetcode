package pathsumii

// https://leetcode.com/problems/path-sum-ii

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	path := []int{}
	dfs(root, targetSum, path, &result)
	return result
}

func dfs(cur *TreeNode, target int, path []int, result *[][]int) {
	if cur == nil {
		return
	}
	if cur.Left != nil {
		path = append(path, cur.Val)
		dfs(cur.Left, target-cur.Val, path, result)
		path = path[:len(path)-1]
	}
	if cur.Right != nil {
		path = append(path, cur.Val)
		dfs(cur.Right, target-cur.Val, path, result)
		path = path[:len(path)-1]
	}
	if cur.Left == nil && cur.Right == nil && target == cur.Val {
		path = append(path, cur.Val)
		copyPath := make([]int, len(path))
		copy(copyPath, path)
		*result = append(*result, copyPath)
		path = path[:len(path)-1]
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
