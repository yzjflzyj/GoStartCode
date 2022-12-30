package main

/**
 * 1109. 航班预订统计
 * 这里有n个航班，它们分别从 1 到 n 进行编号。
 * 有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]意味着
 * 在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。
 * 请你返回一个长度为 n 的数组answer，里面的元素是每个航班预定的座位总数。
 */
//差分数组
func corpFlightBookings(bookings [][]int, n int) []int {
	// 构造差分数组
	counters := make([]int, n)
	for _, v := range bookings {
		counters[v[0]-1] += v[2]
		if v[1] < n {
			counters[v[1]] -= v[2]
		}
	}
	// 差分数组展开
	for i := 1; i < n; i++ {
		counters[i] += counters[i-1]
	}
	return counters
}
