package _9_remove_nth_node_form_end_if_list

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	dummyNode := &types.ListNode{}
	dummyNode.Next = head
	pre := findK(dummyNode, n+1)
	pre.Next = pre.Next.Next
	return dummyNode.Next
}

// findk 找到倒数第k个 单链表要想去除倒数第k个 就得找到倒数第k+1个
func findK(head *types.ListNode, n int) *types.ListNode {
	p1 := head

	for n > 0 {
		p1 = p1.Next
		n--
	}

	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
