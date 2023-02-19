package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *types.ListNode) bool {
	stack := []int{}
	temp := head.Next
	for {
		if len(stack) == 0 {
			stack = append(stack, temp.Val)
		} else {
			if stack[len(stack)-1] == temp.Val {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, temp.Val)
			}
		}

		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func reverseList(head *types.ListNode) *types.ListNode {
	var prev, cur *types.ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *types.ListNode) *types.ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome2(head *types.ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
