package main

/**
 * 1094. 拼车
 * 车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）
 * 给定整数capacity和一个数组 trips , trip[i] = [numPassengersi, fromi, toi]表示第 i 次旅行有numPassengersi乘客，
 * 接他们和放他们的位置分别是fromi和toi。这些位置是从汽车的初始位置向东的公里数。
 * 当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false。
 */
func carPooling(trips [][]int, capacity int) bool {
	// 构建差分数组
	diff := make([]int, 1000)
	for _, v := range trips {
		diff[v[1]] += v[0]
		// 注意，这里v[2]已经下车了
		if v[2] < len(diff) {
			diff[v[2]] -= v[0]
		}
	}
	// 差分数组复原，并检测
	if diff[0] > capacity {
		return false
	}
	arr := make([]int, 1000)
	arr[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		arr[i] = arr[i-1] + diff[i]
		if arr[i] > capacity {
			return false
		}
	}
	return true
}
