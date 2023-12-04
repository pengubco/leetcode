package binary_tree_right_side_view

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	cases := []struct {
		root *TreeNode
		view []int
	}{
		{createTree1(), []int{1, 3, 4}},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.view, rightSideView(tc.root))
		})
	}
}

func createTree1() *TreeNode {
	nodes := make([]TreeNode, 6)
	for i := 1; i <= 5; i++ {
		nodes[i].Val = i
	}
	nodes[1].Left = &nodes[2]
	nodes[1].Right = &nodes[3]
	nodes[2].Right = &nodes[5]
	nodes[3].Right = &nodes[4]
	return &nodes[1]
}
