package reverse_linked_list_ii

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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}
	pre := &ListNode{
		Next: head,
	}
	cur := head
	for i := 0; i < left-1; i++ {
		pre = cur
		cur = cur.Next
	}
	b := pre
	pre = cur
	cur = cur.Next
	for i := 0; i < right-left; i++ {
		n := cur.Next
		cur.Next = pre
		pre = cur
		cur = n
	}
	b.Next.Next = cur
	if left == 1 {
		head = pre
	} else {
		b.Next = pre
	}
	return head
}
