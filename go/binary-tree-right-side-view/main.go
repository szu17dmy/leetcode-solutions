package binary_tree_right_side_view

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

type depth struct {
	node *TreeNode
	depth int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var queue []depth
	queue = append(queue, depth{
		node: root,
		depth: 0,
	})
	var r []int
	for len(queue) != 0 {
		n := queue[0]
		if len(r) > n.depth {
			r[n.depth] = n.node.Val
		} else {
			r = append(r, n.node.Val)
		}
		if n.node.Left != nil {
			queue = append(queue, depth{
				node:  n.node.Left,
				depth: n.depth + 1,
			})
		}
		if n.node.Right != nil {
			queue = append(queue, depth{
				node:  n.node.Right,
				depth: n.depth + 1,
			})
		}
		queue = queue[1:]
	}
	return r
}
