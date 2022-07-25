package _22_coin_change

import "math"

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	if amount == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func coinChange2(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && memo[i] > memo[i-c]+1 {
				memo[i] = memo[i-c] + 1
			}
		}
	}
	if memo[amount] == math.MaxInt32 {
		return -1
	}
	return memo[amount]
}
