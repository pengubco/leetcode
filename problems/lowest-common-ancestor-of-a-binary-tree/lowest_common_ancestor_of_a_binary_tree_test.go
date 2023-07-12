package lowest_common_ancestor_of_a_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor_1(t *testing.T) {
	nodes := make([]*TreeNode, 9)
	for i := 0; i < 9; i++ {
		nodes[i] = &TreeNode{
			Val:   i,
			Left:  nil,
			Right: nil,
		}
	}
	nodes[3].Left, nodes[3].Right = nodes[5], nodes[1]
	nodes[5].Left, nodes[5].Right = nodes[6], nodes[2]
	nodes[2].Left, nodes[2].Right = nodes[7], nodes[4]
	nodes[1].Left, nodes[1].Right = nodes[0], nodes[8]

	assert.Equal(t, nodes[3], lowestCommonAncestor(nodes[3], nodes[5], nodes[1]))
	assert.Equal(t, nodes[5], lowestCommonAncestor(nodes[3], nodes[5], nodes[4]))
}
