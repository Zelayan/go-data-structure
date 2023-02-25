package leetcode

// 暴力解法
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 这里的dp定义是前i天卖出股票所能获取的最大利润
	dp := make([]int, len(prices))
	dp[0] = 0
	// 用来记录最小的元素
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		// 这里比较从之前的最低价格买，到今天卖的价格高，还是
		// 之前的价格高，
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return dp[len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
