package path_sum

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

func hasPathSum(root *TreeNode, targetSum int) bool {
	var stack []*TreeNode
	cur := root
	var prev *TreeNode
	var sum int
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			sum += cur.Val
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		if cur.Left == nil && cur.Right == nil && sum == targetSum {
			return true
		}
		if cur.Right == nil || cur.Right == prev {
			stack = stack[:len(stack)-1]
			sum -= cur.Val
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return false
}
