package populating_next_right_pointers_in_each_node_ii

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			n := q[i]
			if i < l-1 {
				n.Next = q[i+1]
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		q = q[l:]
	}
	return root
}
