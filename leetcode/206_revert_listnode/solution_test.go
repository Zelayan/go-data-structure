package _06_revert_listnode

import (
	"github.com/Zelayan/go-fucking-written-examination/types"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "success",
			args: args{head: &types.ListNode{
				Val: 1,
				Next: &types.ListNode{
					Val: 2,
					Next: &types.ListNode{
						Val:  3,
						Next: nil,
					},
				},
			}},
			want: &types.ListNode{
				Val: 3,
				Next: &types.ListNode{
					Val: 2,
					Next: &types.ListNode{
						Val:  1,
						Next: nil,
					},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
