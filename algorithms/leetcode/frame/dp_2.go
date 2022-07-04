package main

//322. 零钱兑换
//动态规划
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			//跳过不符合的
			if i-coin < 0 {
				continue
			}
			//取更小值
			if dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
