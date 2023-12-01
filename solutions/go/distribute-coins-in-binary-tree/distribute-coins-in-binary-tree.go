package distributecoinsinbinarytree

func distributeCoins(root *TreeNode) int {
	var ans int
	dfs(root, &ans)
	return ans
}

// return size of the subtree, and the total coins of the subtree.
func dfs(cur *TreeNode, ans *int) (int, int) {
	if cur == nil {
		return 0, 0
	}
	lSize, lCoins := dfs(cur.Left, ans)
	rSize, rCoins := dfs(cur.Right, ans)
	size := lSize + rSize + 1
	coins := cur.Val + lCoins + rCoins
	if coins > size {
		*ans += coins - size
	} else {
		*ans += size - coins
	}
	return size, coins
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
