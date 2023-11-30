package average_of_levels_in_binary_tree

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

func averageOfLevels(root *TreeNode) []float64 {
	var queue []*TreeNode
	queue = append(queue, root)
	var r []float64
	for len(queue) != 0 {
		var sum int
		l := len(queue)
		for i := 0; i < l; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		r = append(r, float64(sum)/float64(l))
		queue = queue[l:]
	}
	return r
}
