package populatingnextrightpointersineachnodeii

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	nodes := make([]Node, 8)
	for i := 1; i <= 7; i++ {
		nodes[i].Val = i
		if i <= 3 {
			nodes[i].Left = &nodes[2*i]
			nodes[i].Right = &nodes[2*i+1]
		}
	}
	nodes[3].Left = nil

	connect(&nodes[1])
	for i := 1; i <= 7; i++ {
		fmt.Printf("%d %v\n", nodes[i].Val, nodes[i].Next)
	}
}
