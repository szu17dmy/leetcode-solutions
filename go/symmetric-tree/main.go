package symmetric_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return validate(root.Left, root.Right)
}

func validate(lhs, rhs *TreeNode) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}
	return lhs.Val == rhs.Val && validate(lhs.Left, rhs.Right) && validate(lhs.Right, rhs.Left)
}
