package binarytree

import (
	"reflect"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	type args struct {
		i    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *BinaryTree
	}{
		{
			name: "success",
			args: args{
				i:    0,
				nums: []int{1, 2, 3},
			},
			want: &BinaryTree{Value: 1, Left: &BinaryTree{Value: 2}, Right: &BinaryTree{Value: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateBinaryTree(tt.args.i, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
