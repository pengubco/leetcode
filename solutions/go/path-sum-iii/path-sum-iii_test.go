package pathsumiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum_1(t *testing.T) {
	assert.Equal(t, 3, pathSum(tree1(), 8))
}

func tree1() *TreeNode {
	nodes := make([]TreeNode, 9)
	nodes[0] = TreeNode{Val: 10, Left: &nodes[1], Right: &nodes[2]}
	nodes[1] = TreeNode{Val: 5, Left: &nodes[3], Right: &nodes[4]}
	nodes[2] = TreeNode{Val: -3, Right: &nodes[5]}
	nodes[3] = TreeNode{Val: 3, Left: &nodes[6], Right: &nodes[7]}
	nodes[4] = TreeNode{Val: 2, Right: &nodes[8]}
	nodes[5] = TreeNode{Val: 11}
	nodes[6] = TreeNode{Val: 3}
	nodes[7] = TreeNode{Val: -2}
	nodes[8] = TreeNode{Val: 1}
	return &nodes[0]
}
