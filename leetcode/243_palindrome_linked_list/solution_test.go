package leetcode

import (
	"github.com/Zelayan/go-fucking-written-examination/types"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{head: types.GenListNode([]int{1, 2, 2, 1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{head: types.GenListNode([]int{1, 2, 2, 1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
