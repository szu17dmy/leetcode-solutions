package binary_tree_level_order_traversal

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

func levelOrder(root *TreeNode) (r [][]int) {
	if root == nil {
		return
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var level int
	for len(queue) != 0 {
		r = append(r, []int{})
		l := len(queue)
		for i := 0; i < l; i++ {
			r[level] = append(r[level], queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
		level++
	}
	return
}
