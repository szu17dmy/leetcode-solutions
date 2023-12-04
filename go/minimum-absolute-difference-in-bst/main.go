package minimum_absolute_difference_in_bst

import "math"

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

func getMinimumDifference(root *TreeNode) int {
	var stack []*TreeNode
	cur := root
	prev := cur.Val
	delta := math.MaxInt
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		d := cur.Val - prev
		if d > 0 && d < delta {
			delta = d
		}
		prev = cur.Val
		cur = cur.Right
	}
	return delta
}
