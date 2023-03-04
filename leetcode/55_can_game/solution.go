package leetcode

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	// dp[i] 代表可以走到的最大距离
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if dp[0] == 0 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if dp[i-1] == i && nums[i] == 0 {
			if dp[i-1] >= len(nums)-1 {
				return true
			}
			return false
		}
		dp[i] = max(dp[i-1], i+nums[i])
		if dp[i] >= len(nums)-1 {
			return true
		}
	}

	return dp[len(nums)-1] >= len(nums)
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
