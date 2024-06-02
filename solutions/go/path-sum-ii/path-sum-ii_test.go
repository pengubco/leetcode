package pathsumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	tree1Nodes := make([]TreeNode, 10)
	tree1Nodes[0] = TreeNode{
		Val:   5,
		Left:  &tree1Nodes[1],
		Right: &tree1Nodes[2],
	}
	tree1Nodes[1] = TreeNode{
		Val:  4,
		Left: &tree1Nodes[3],
	}
	tree1Nodes[2] = TreeNode{
		Val:   8,
		Left:  &tree1Nodes[4],
		Right: &tree1Nodes[5],
	}
	tree1Nodes[3] = TreeNode{
		Val:   11,
		Left:  &tree1Nodes[6],
		Right: &tree1Nodes[7],
	}
	tree1Nodes[4] = TreeNode{
		Val: 13,
	}
	tree1Nodes[5] = TreeNode{
		Val:   4,
		Left:  &tree1Nodes[8],
		Right: &tree1Nodes[9],
	}
	tree1Nodes[6] = TreeNode{Val: 7}
	tree1Nodes[7] = TreeNode{Val: 2}
	tree1Nodes[8] = TreeNode{Val: 5}
	tree1Nodes[9] = TreeNode{Val: 1}

	result := pathSum(&tree1Nodes[0], 22)
	assert.Equal(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, result)
}
