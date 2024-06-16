package binarytreecameras

// https://leetcode.com/problems/binary-tree-cameras/

func minCameraCover(root *TreeNode) int {
	tc := &TreeCamera{
		root:    root,
		covered: make(map[*TreeNode]bool),
	}
	tc.MinCameraCover(root, nil)
	return tc.NumberOfCameras()
}

type TreeCamera struct {
	root            *TreeNode
	covered         map[*TreeNode]bool
	numberOfCameras int
}

func (tc *TreeCamera) MinCameraCover(cur *TreeNode, parent *TreeNode) {
	children := []*TreeNode{}
	if cur.Left != nil {
		children = append(children, cur.Left)
	}
	if cur.Right != nil {
		children = append(children, cur.Right)
	}
	for _, child := range children {
		tc.MinCameraCover(child, cur)
	}
	needsToCoverChildren := false
	for _, children := range children {
		if !tc.covered[children] {
			needsToCoverChildren = true
			break
		}
	}
	// If cur has some children to cover, put a camera at cur.
	// If cur is the root and uncovered, put a camera.
	if needsToCoverChildren || (parent == nil && !tc.covered[cur]) {
		tc.numberOfCameras++
		tc.covered[cur] = true
		if parent != nil {
			tc.covered[parent] = true
		}
		for _, children := range children {
			tc.covered[children] = true
		}
	}
}

func (tc *TreeCamera) NumberOfCameras() int {
	return tc.numberOfCameras
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
