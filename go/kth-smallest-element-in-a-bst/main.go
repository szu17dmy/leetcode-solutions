package kth_smallest_element_in_a_bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	node, cnt := root, k
	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		cnt--
		if cnt == 0 {
			return node.Val
		}
		node = node.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
