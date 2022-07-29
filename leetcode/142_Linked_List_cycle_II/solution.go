package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 通过map 处理，走过的路记录一下， 若再次走过则时环的起点
func detectCycle(head *types.ListNode) *types.ListNode {
	m := make(map[*types.ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}
