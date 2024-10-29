package main

import "fmt"

type Relation struct {
	Child  int
	Parent int // -1 if don't have
	IsLeft bool
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
15 20 true
19 80 true
17 20 false
16 80 false
80 50 false
50 -1 false
20 50 true
*
*/
func BuildBinaryTree(relations []Relation) *TreeNode {
	// find the root
	root := &TreeNode{}
	nodeIdx := make(map[int][]Relation)

	for _, relation := range relations {
		if relation.Parent == -1 {
			root.Val = relation.Child
		} else {
			nodeIdx[relation.Parent] = append(nodeIdx[relation.Parent], relation)
		}
	}

	buildTree(root, nodeIdx)
	return root
}

func inOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	inOrderTraversal(node.Left)
	fmt.Printf("%d ", node.Val)
	inOrderTraversal(node.Right)
}

func buildTree(node *TreeNode, idx map[int][]Relation) {
	if _, exists := idx[node.Val]; !exists {
		return
	}

	relations := idx[node.Val]
	for _, relation := range relations {
		newNode := &TreeNode{Val: relation.Child}
		if relation.IsLeft {
			node.Left = newNode
		} else {
			node.Right = newNode
		}
	}

	buildTree(node.Left, idx)
	buildTree(node.Right, idx)
}

func main() {
	relations := []Relation{
		{15, 20, true},
		{19, 80, true},
		{17, 20, false},
		{16, 80, false},
		{80, 50, false},
		{50, -1, false},
		{20, 50, true},
	}

	root := BuildBinaryTree(relations)
	fmt.Println("Binary Tree:")
	inOrderTraversal(root)
	fmt.Println()
}
