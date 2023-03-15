package leetcode

import (
	"testing"

	"github.com/Zelayan/go-fucking-written-examination/types"
)

func Test_findTilt(t *testing.T) {
	type args struct {
		root *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				root: types.BuildTree([]int{4,2,9,3,5,types.TreeNull,7}),
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
