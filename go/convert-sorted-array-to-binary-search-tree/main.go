package convert_sorted_array_to_binary_search_tree

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

func sortedArrayToBST(nums []int) *TreeNode {
	var build func([]int, int, int) *TreeNode
	build = func(nums []int, begin, end int) *TreeNode {
		if begin > end {
			return nil
		}
		mid := (begin + end + 1) / 2
		node := &TreeNode{
			Val: nums[mid],
		}
		node.Left = build(nums, begin, mid-1)
		node.Right = build(nums, mid+1, end)
		return node
	}
	return build(nums, 0, len(nums)-1)
}
