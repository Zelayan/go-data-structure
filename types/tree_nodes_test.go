package types

import (
	"fmt"
	"reflect"
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

func TestGenListNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "success",
			args: args{[]int{1, 2, 3}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "success",
			args: args{[]int{22, 2, 33}},
			want: &ListNode{Val: 22, Next: &ListNode{Val: 2, Next: &ListNode{Val: 33}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
