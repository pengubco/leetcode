package pathsumiii

func pathSum(root *TreeNode, targetSum int) int {
	sumCount := make(map[int]int)
	sumCount[0] = 1
	targetCount := 0
	dfs(root, targetSum, 0, sumCount, &targetCount)
	return targetCount
}

func dfs(cur *TreeNode, targetSum int, pathSum int, sumCount map[int]int, targetCount *int) {
	if cur == nil {
		return
	}
	pathSum += cur.Val
	*targetCount += sumCount[pathSum-targetSum]
	sumCount[pathSum]++
	dfs(cur.Left, targetSum, pathSum, sumCount, targetCount)
	dfs(cur.Right, targetSum, pathSum, sumCount, targetCount)
	sumCount[pathSum]--
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
