package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return check(root1, root2)
}

func check(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}

	return (check(node1.Left, node2.Left) || check(node1.Left, node2.Right)) && (check(node1.Right, node2.Left) || check(node1.Right, node2.Right))
}

func main() {
}
