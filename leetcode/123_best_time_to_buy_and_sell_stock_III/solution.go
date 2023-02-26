package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	// 任意一天结束之后，我们会处于
	// 1. 未进行操作
	// 2. 进行了买
	// 3. 进行一次买一次卖，完成一次交易
	// 4. 完成一次交易的情况下，进行第二次买
	// 5. 完成了两次交易
	//
	//

	// 分别用 buy1、sell1、buy2和sell2 表示 2,3,4,5的最大利润

	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0

	// 1. 利润为0
	// 2. 利润为 前一天买的利润

	for i := 1; i < n; i++ {
		// 第一次买，最大利润是 上次买的利润，当前买的话利润就得减去当前的价格
		buy1 = max(buy1, -prices[i])
		// 第一次卖，利润是买之前的利润，加上当天卖出去的价格
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
