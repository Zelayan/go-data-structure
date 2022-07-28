package _76_middle_if_list_node

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 利用快慢指针，当（偶数个节点）当快指针不够一步时，这时候满指针刚还在中间节点的第一个，快慢指针继续往下走一步，慢指针刚好走到第二个节点
func middleNode(head *types.ListNode) *types.ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow

}
