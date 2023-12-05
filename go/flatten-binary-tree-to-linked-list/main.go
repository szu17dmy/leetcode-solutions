package flatten_binary_tree_to_linked_list

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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var stack, helper []*TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			helper = append(helper, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	for i := 0; i < len(helper)-1; i++ {
		helper[i].Left = nil
		helper[i].Right = helper[i+1]
	}
	helper[len(helper)-1].Left = nil
	helper[len(helper)-1].Right = nil
}
