package leetcode

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}
	// 首先要完成k笔交易
	len := len(prices)

	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < k; i++ {
		dp[i][0] = -prices[0]
		dp[i][1] = 0
	}

	for i := 0; i < len; i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				// 如果是第一次笔
				dp[j][0] = max(dp[j][0], -prices[i])
				dp[j][1] = max(dp[j][1], dp[j][0]+prices[i])
			} else {
				//
				dp[j][0] = max(dp[j][0], dp[j-1][1]-prices[i])
				//
				dp[j][1] = max(dp[j][1], dp[j][0]+prices[i])
			}
		}
	}
	// 最后一笔卖出的利润最高
	return dp[k-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
