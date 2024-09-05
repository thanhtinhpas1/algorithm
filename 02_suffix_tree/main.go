package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Label []byte
	Nodes []*Node
}

func (n *Node) String() string {
	var str = fmt.Sprintf("%s\n", string(n.Label))

	for _, n := range n.Nodes {
		str = fmt.Sprintf("%s%s", str, n.String())
	}

	return str
}

func (node *Node) insertNode(edge *Node) {
	newNodeLen := len(edge.Label)
	
	idx := sort.Search(len(node.Nodes), func(i int) bool {
		return newNodeLen < len(node.Nodes[i].Label)
	})

	node.Nodes = append(node.Nodes, nil)
	copy(node.Nodes[idx+1:], node.Nodes[idx:])
	
	node.Nodes[idx] = edge
}

func (node *Node) removeNode(idx int) {
	copy(node.Nodes[idx:], node.Nodes[idx+1:])
	node.Nodes[len(node.Nodes) - 1] = nil
	node.Nodes = node.Nodes[:len(node.Nodes)-1]
}

func main() {
	node3 := Node{
		Label: []byte("abcd"),
		Nodes: nil,
	}

	node2 := Node{
		Label: []byte("ab"),
		Nodes: nil,
	}

	node1 := Node{
		Label: []byte("a"),
		Nodes: []*Node{&node2, &node3},
	}
	fmt.Println(node1.String())


	node1.insertNode(&Node{
		Label: []byte("abc"),
		Nodes: nil,
	})

	fmt.Println(node1.String())
}