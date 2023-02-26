package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	res := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			res = res + prices[i] - prices[i-1]
		}
	}
	return res
}
