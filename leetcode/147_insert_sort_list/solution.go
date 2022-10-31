package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertionSortList(head *types.ListNode) *types.ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &types.ListNode{Next: head}
	// lastSorted
    lastSorted, curr := head, head.Next
    for curr != nil {
        if lastSorted.Val < curr.Val {
            lastSorted = curr
        } else {
            prev := dummyHead
            for prev.Next.Val < curr.Val {
                prev = prev.Next
            }
            lastSorted.Next = curr.Next
            curr.Next = prev.Next
            prev.Next = curr
        }
        curr = lastSorted.Next
    }
    return dummyHead.Next
}