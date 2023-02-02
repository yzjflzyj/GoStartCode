package main

/**
 * 53. 最大子数组和
 */
//动态规划:dp[i]表示以i索引元素结尾的,最大连续数组和
//dp[i]仅与dp[i-1]有关,因此可以迭代优化
func maxSubArray(nums []int) int {
	dp := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp += nums[i]
		} else {
			dp = nums[i]
		}
		//res = int(math.Max(float64(res), float64(dp)))
		res = max(res, dp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
