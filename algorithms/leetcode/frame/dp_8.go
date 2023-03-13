package main

/**
 * 1143. 最长公共子序列
 */
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxValue(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
func maxValue(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
