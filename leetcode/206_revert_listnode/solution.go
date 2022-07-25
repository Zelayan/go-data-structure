package _06_revert_listnode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return last
}
