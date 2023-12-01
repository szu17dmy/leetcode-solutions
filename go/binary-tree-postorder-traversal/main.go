package binary_tree_postorder_traversal

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

func postorderTraversal(root *TreeNode) (r []int) {
	var stack []*TreeNode
	node := root
	var prev *TreeNode
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == prev {
			stack = stack[:len(stack)-1]
			r = append(r, node.Val)
			prev = node
			node = nil
		} else {
			node = node.Right
		}
	}
	return r
}
