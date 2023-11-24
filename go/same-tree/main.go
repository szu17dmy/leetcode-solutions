package same_tree

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	var qp, qq []*TreeNode
	qp = append(qp, p)
	qq = append(qq, q)
	for len(qp) != 0 {
		if qp[0].Left != nil {
			if qq[0].Left != nil {
				qp = append(qp, qp[0].Left)
				qq = append(qq, qq[0].Left)
			} else {
				return false
			}
		} else {
			if qq[0].Left != nil {
				return false
			}
		}
		if qp[0].Right != nil {
			if qq[0].Right != nil {
				qp = append(qp, qp[0].Right)
				qq = append(qq, qq[0].Right)
			} else {
				return false
			}
		} else {
			if qq[0].Right != nil {
				return false
			}
		}
		if qp[0].Val != qq[0].Val {
			return false
		}
		qp = qp[1:]
		qq = qq[1:]
	}
	return true
}
