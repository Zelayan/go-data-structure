package _30_K_Smallest_Element_in_BST

import (
	"github.com/Zelayan/go-fucking-written-examination/types"
	"testing"
)

func TestKsmElementInBST(t *testing.T) {
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
			args: args{
				root: &types.TreeNode{
					Val: 5,
					Left: &types.TreeNode{
						Val: 3,
						Left: &types.TreeNode{
							Val: 2,
							Left: &types.TreeNode{
								Val: 1,
							},
						},
						Right: &types.TreeNode{Val: 4},
					},
					Right: &types.TreeNode{
						Val: 6,
					},
				},
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KsmElementInBST(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("KsmElementInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
