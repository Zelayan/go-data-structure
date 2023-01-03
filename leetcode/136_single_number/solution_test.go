package leetcode

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
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
			args: args{nums: []int{1, 1, 2, 2, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	// 1. 任何数和0异或，都是其本身 a^0 = a
	// 2. 任何数和其本身异或，都等于0 a^a = 0
	// 3. 异或运算满足交换律和结合律，a ^ b ^ a = a ^ a ^b = 0 ^ b = b
	a := 0
	fmt.Println(a ^ 10)
	fmt.Println(10 ^ 10)
}
