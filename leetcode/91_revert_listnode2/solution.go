package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseBetween(head *types.ListNode, left int, right int) *types.ListNode {
    dummyNode := &types.ListNode{Val: -1}
    dummyNode.Next = head
    pre := dummyNode
    for i := 0; i < left -1; i++ {
        pre = pre.Next
    }
    cur := pre.Next
	// 采用头插法，反转链表
	// pre永远指向新的反转后的header
	// cur指向待反转的第一个节点
	// next永远指向 cur的next
    for i := 0; i < right-left; i++ {
        next := cur.Next   
        cur.Next =cur.Next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return dummyNode.Next
}