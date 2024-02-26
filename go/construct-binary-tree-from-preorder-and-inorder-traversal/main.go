package construct_binary_tree_from_preorder_and_inorder_traversal

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{
		Val: preorder[0],
	}
	stack := []*TreeNode{root}
	for i, j := 1, 0; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		if node.Val != inorder[j] {
			node.Left = &TreeNode{
				Val: preorder[i],
			}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[j] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				j++
			}
			node.Right = &TreeNode{
				Val: preorder[i],
			}
			stack = append(stack, node.Right)
		}
	}
	return root
}
