package _54_maximum_binary_tree

import (
	"github.com/Zelayan/go-fucking-written-examination/types"
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *types.TreeNode
	}{
		{
			name: "success",
			args: args{nums: []int{3, 2, 1, 6, 0, 5}},
			want: &types.TreeNode{
				Val: 6,
				Left: &types.TreeNode{
					Val:  3,
					Left: nil,
					Right: &types.TreeNode{
						Val:  2,
						Left: nil,
						Right: &types.TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &types.TreeNode{
					Val: 5,
					Left: &types.TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xx(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	t.Log(a[:0])
	t.Log(a[2:])
}
