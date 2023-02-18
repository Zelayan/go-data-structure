package leetcode

import "math"

// 整数拆分
func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	if n%3 == 0 {
		return int(math.Floor(math.Pow(3, float64(n/3))))
	}

	if n%3 == 1 {
		return int(math.Pow(3, math.Floor(float64(n/3))-1) * 2 * 2)
	}
	if n%3 == 2 {
		return int(math.Pow(3, math.Floor(float64(n/3))) * 2)
	}
	return 0
}
