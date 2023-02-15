package leetcode

import "testing"

func Test_getArea(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{nums: []int{3, 1, 2, 3}},
			want: 3,
		},
		{
			name: "success",
			args: args{nums: []int{0, 1, 2, 3}},
			want: 0,
		},
		{
			name: "succe2ss",
			args: args{nums: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getArea(tt.args.nums); got != tt.want {
				t.Errorf("getArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{nums: []int{1, 0, 1}},
			want: 1,
		},

		{
			name: "succe2ss",
			args: args{nums: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.nums); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
