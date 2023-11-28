package merge_two_sorted_lists

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	root := &ListNode{}
	ptr := root
	for p1 != nil || p2 != nil {
		if p1 != nil && p2 != nil {
			if p1.Val > p2.Val {
				ptr.Next = &ListNode{
					Val: p2.Val,
				}
				p2 = p2.Next
			} else {
				ptr.Next = &ListNode{
					Val: p1.Val,
				}
				p1 = p1.Next
			}
		} else if p1 != nil {
			ptr.Next = &ListNode{
				Val: p1.Val,
			}
			p1 = p1.Next
		} else {
			ptr.Next = &ListNode{
				Val: p2.Val,
			}
			p2 = p2.Next
		}
		ptr = ptr.Next
	}
	return root.Next
}
