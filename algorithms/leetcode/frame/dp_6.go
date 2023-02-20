package main

import "sort"

/**
 *354. 俄罗斯套娃信封问题
 *给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
 *当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *注意：不允许旋转信封。
 */
func maxEnvelopes(envelopes [][]int) int {
	//先按长度升序，再按宽度降序
	sort.Slice(len(envelopes), func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	// dp表示以envelopes的当前元素结尾的最长宽度递增子序列
	var dp []int
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(dp, h); i < len(dp) {
			dp[i] = h
		} else {
			dp = append(dp, h)
		}
	}
	return len(dp)
}

func main() {

}
