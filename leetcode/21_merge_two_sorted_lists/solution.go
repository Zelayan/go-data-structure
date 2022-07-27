package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	dummyNode := &types.ListNode{}
	p := dummyNode
	l := list1
	r := list2

	for l != nil && r != nil {
		if l.Val > r.Val {
			p.Next = r
			r = r.Next
		} else {
			p.Next = l
			l = l.Next
		}

		p = p.Next
	}
	if l != nil {
		p.Next = l
	}
	if r != nil {
		p.Next = r
	}

	return dummyNode.Next
}
