package sum_root_to_leaf_numbers

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum = 10*sum + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return dfs(node.Left, sum) + dfs(node.Right, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
