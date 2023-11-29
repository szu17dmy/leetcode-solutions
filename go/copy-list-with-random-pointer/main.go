package copy_list_with_random_pointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cache := map[*Node]*Node{}
	p1 := head
	r := &Node{}
	p2 := r
	for p1 != nil {
		p2.Next = &Node{}
		p2 = p2.Next
		p2.Val = p1.Val
		if _, ok := cache[p1]; !ok {
			cache[p1] = p2
		}
		p1 = p1.Next
	}
	p1, p2 = head, r.Next
	for p1 != nil {
		p2.Random = cache[p1.Random]
		p1 = p1.Next
		p2 = p2.Next
	}
	return r.Next
}
