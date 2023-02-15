package leetcode

import "math"

func findUnsortedSubarray(nums []int) int {
	ret := 0
	left, right := -1, -1
	min, max := math.MaxInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		// 找到最大的数据
		if nums[i] >= max {
			max = nums[i]
		} else {
			//如果后面比最大的数据还小，则说明需要交换
			right = i
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		// 找到最新的数据
		if nums[i] <= min {
			min = nums[i]
		} else {
			//如果后面的数据比最小的海啸吗，说明需要交换
			left = i
		}
	}
	if right > left {
		ret = right - left + 1
	}
	return ret
}
