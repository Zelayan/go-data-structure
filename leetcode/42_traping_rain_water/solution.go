package leetcode

// 备忘录模式
// nums[i] 位置能接多少雨水，取决于它的左边最高，和右边最高中最低的
func trap(nums []int) int {
	res := 0
	n := len(nums)
	// 如果小于3 说明 接不到雨水
	if n < 3 {
		return 0
	}
	lMax := []int{nums[0]}
	rMax := make([]int, n)
	rMax[n-1] = nums[n-1]

	// 每个位置的左边最大计算出来
	for i := 1; i < n; i++ {
		lMax = append(lMax, max(nums[i], lMax[i-1]))
	}

	// 再把每个位置右边最大计算出来
	for i := n - 2; i > -1; i-- {
		rMax[i] = max(nums[i], rMax[i+1])
	}

	// 从1开始，是因为下标为0的接不到雨水
	for i := 1; i < n; i++ {
		res = res + min(lMax[i], rMax[i]) - nums[i]
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
