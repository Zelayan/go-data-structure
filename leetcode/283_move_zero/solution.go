package leetcode

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && nums[zero] != 0 {
			zero = i
		}
		if nums[i] != 0 && nums[zero] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			// 这里需要确定从左到右第一个零的位置 用于交换
			for j := zero + 1; j <= i; j++ {
				if nums[j] == 0 {
					zero = j
					break
				}
			}
		}
	}
}
