package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				nums1: []int{1,2,2,3},
				m : 4,
				nums2: []int{2,2},
				n : 2,
			},
			want: []int{1,2,2,2,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
