package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 利用快慢指针，快指针走一步走两格，慢指针一步走一格，如果有环，快指针一定会追上慢指针
func hasCycle(head *types.ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
