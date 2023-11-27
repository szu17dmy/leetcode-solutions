package add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	ptr := root
	var carried int
	p1, p2 := l1, l2
	for p1 != nil || p2 != nil {
		if p1 != nil && p2 != nil {
			ptr.Next = &ListNode{
				Val: (p1.Val + p2.Val + carried) % 10,
			}
			carried = (p1.Val + p2.Val + carried) / 10
			p1 = p1.Next
			p2 = p2.Next
		} else if p1 != nil {
			ptr.Next = &ListNode{
				Val: (p1.Val + carried) % 10,
			}
			carried = (p1.Val + carried) / 10
			p1 = p1.Next
		} else {
			ptr.Next = &ListNode{
				Val: (p2.Val + carried) % 10,
			}
			carried = (p2.Val + carried) / 10
			p2 = p2.Next
		}
		ptr = ptr.Next
	}
	if carried != 0 {
		ptr.Next = &ListNode{
			Val: 1,
		}
	}
	return root.Next
}
