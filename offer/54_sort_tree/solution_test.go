package leetcode

import (
	"testing"

	"github.com/Zelayan/go-fucking-written-examination/types"
)

func Test_kthLargest(t *testing.T) {
	type args struct {
		root *types.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{ root:types.BuildTree([]int{3,1,4,types.TreeNull,2}),
					k:1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
