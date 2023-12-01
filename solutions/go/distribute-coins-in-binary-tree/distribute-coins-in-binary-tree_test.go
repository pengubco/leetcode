package distributecoinsinbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCoins(t *testing.T) {
	cases := []struct {
		root *TreeNode
		ans  int
	}{
		{createTree1(), 2},
		{createTree2(), 3},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.ans, distributeCoins(tc.root))
		})
	}
}

func createTree1() *TreeNode {
	nodes := []TreeNode{
		TreeNode{Val: 3},
		TreeNode{},
		TreeNode{},
	}
	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]
	return &nodes[0]
}

func createTree2() *TreeNode {
	nodes := []TreeNode{
		TreeNode{},
		TreeNode{Val: 3},
		TreeNode{},
	}
	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]
	return &nodes[0]
}
