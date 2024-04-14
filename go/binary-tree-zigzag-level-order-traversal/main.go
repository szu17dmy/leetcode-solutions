package binary_tree_zigzag_level_order_traversal

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

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	var level int
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		ans = append(ans, []int{})
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			idx := i
			if level%2 != 0 {
				idx = l-1-i
			}
			ans[level] = append(ans[level], queue[idx].Val)
		}
		queue = queue[l:]
		level++
	}
	return ans
}
