package binarytreecameras

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCameraCover(t *testing.T) {
	cases := []struct {
		root   *TreeNode
		result int
	}{
		{
			root:   createTree1(),
			result: 1,
		},
		{
			root:   createTree2(),
			result: 1,
		},
		{
			root:   createTree3(),
			result: 2,
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.result, minCameraCover(tc.root))
		})
	}
}

func createTree2() *TreeNode {
	/*
	   0
	*/
	nodes := make([]TreeNode, 1)
	return &nodes[0]
}

func createTree1() *TreeNode {
	/*
		   0
		   1
		2	    3
	*/
	nodes := make([]TreeNode, 4)
	for i := 0; i < 4; i++ {
		nodes[i].Val = i
	}
	nodes[1].Left = &nodes[2]
	nodes[1].Right = &nodes[3]
	nodes[0].Left = &nodes[1]
	return &nodes[0]
}

func createTree3() *TreeNode {
	nodes := make([]TreeNode, 4)
	for i := 0; i < 4; i++ {
		nodes[i].Val = i
	}
	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]
	nodes[2].Right = &nodes[3]
	return &nodes[0]
}
