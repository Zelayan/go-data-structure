package leetcode

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// dp的定义是 dp[i] 表示当前的最长的递增子序列的个数
	dp := make([]int, len(nums))
	dp[0] = 1  // 默认值为1
	maxes := 1 // 用于存放最大值
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 默认为1
		// 从之前到现在遍历
		for j := 0; j < i; j++ {
			// 只要比当前数字小，就说明最长子序列可以加1，
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 去获取最大值
		maxes = max(maxes, dp[i])
	}

	return maxes
}

func max(maxes int, i int) int {
	if maxes > i {
		return maxes
	} else {
		return i
	}
}
