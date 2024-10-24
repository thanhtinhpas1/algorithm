package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	dfs([]*TreeNode{root})
}

func dfs(arr []*TreeNode) {
	if len(arr) == 0 {
		return
	}

	sum := 0

	// calculate sum of all childrens
	for _, node := range arr {
		if node == nil {
			continue
		}

		if node.Left != nil {
			sum += node.Left.Val
		}

		if node.Right != nil {
			sum += node.Right.Val
		}
	}

	var childr []*TreeNode

	for _, node := range arr {
		curSum := 0

		if node.Left != nil {
			curSum += node.Left.Val
		}

		if node.Right != nil {
			curSum += node.Right.Val
		}

		if node.Left != nil {
			node.Left.Val = sum - curSum
			childr = append(childr, node.Left)
		}

		if node.Right != nil {
			node.Right.Val = sum - curSum
			childr = append(childr, node.Right)
		}
	}

	dfs(childr)
}

func main() {

}
