package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *types.ListNode, x int) *types.ListNode {
	dummy1 := &types.ListNode{} // 虚拟节点记录比x小的
	dummy2 := &types.ListNode{} // 虚拟节点记录比相等和大的

	d1 := dummy1
	d2 := dummy2
	for head != nil {
		if head.Val >= x {
			d2.Next = head
			d2 = d2.Next
		} else {
			d1.Next = head
			d1 = d1.Next
		}
		head = head.Next
	}

	d1.Next = dummy2.Next // 将大的拼到小的后面
	d2.Next = nil         //将后续置空
	return dummy1.Next
}
