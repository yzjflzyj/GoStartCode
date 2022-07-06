package main

//518. 零钱兑换 II
func change(amount int, coins []int) int {
	//dp定义:dp[i]表示i金额
	dp := make([]int, amount+1)
	// base case
	dp[0] = 1
	// 遍历顺序
	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			// 推导公式
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
