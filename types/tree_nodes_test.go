package types

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {

	head := &ListNode{Val: 1, Next: &ListNode{
		Val:  2,
		Next: nil,
	}}

	a := head
	fmt.Printf("head: %p, value %v\n", head, head) // 地址一样
	fmt.Printf("a   : %p, value %v\n", a, a)       // 地址一样

	a.Next = &ListNode{Val: 1, Next: &ListNode{
		Val:  2,
		Next: nil,
	}}
	fmt.Printf("head: %p, value %v\n", head, head) // 地址一样
	fmt.Printf("a   : %p, value %v\n", a, a)
	a = a.Next

	fmt.Printf("head: %p, value %v\n", head, head) // 地址一样
	fmt.Printf("a   : %p, value %v\n", a, a)
	for head != nil {
		t.Log(head)
		head = head.Next
	}
}
