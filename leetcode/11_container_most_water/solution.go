package leetcode

func maxArea(height []int) int {

	l, r := 0, len(height)-1
	m := 0
	for l < r {
		m = Max(m, (r-l)*Min(height[l], height[r]))
		// 每次找最短的去替换
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return m
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
