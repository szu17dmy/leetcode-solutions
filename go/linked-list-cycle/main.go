package linked_list_cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		if slow != nil {
			slow = slow.Next
		}
		if fast != nil {
			fast = fast.Next
			if fast != nil {
				fast = fast.Next
			}
		}
	}
	return fast != nil
}
