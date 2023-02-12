package leetcode

// 暴力解法 此方法在leetcode上会超出时间限制
func dailyTemperatures(temperatures []int) []int {
	var res []int
	for i := 0; i < len(temperatures)-1; i++ {
		// 从后续找
		j := i + 1
		for ; j < len(temperatures); j++ {
			// 如果更高 说明找到了
			if temperatures[j] > temperatures[i] {
				res = append(res, j-i)
				break
			}
		}
		// 如果到最后都没有找到，说明是0
		if j == len(temperatures) {
			res = append(res, 0)
		}
	}
	// 最后一个一定是0
	return append(res, 0)
}

// 单调栈解法
func dailyTemperatures2(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int

	for i, v := range temperatures {
		// 如果栈不为空，并且当前这个数，比栈顶大，说明当前这个数就是 当前栈顶的 下一个大的数
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			// 就让栈顶 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 说明当前这个已经找到
			res[top] = i - top
		}
		// 将当前要找的数，压到栈里面，等下次遍历
		stack = append(stack, i)
	}
	return res
}
