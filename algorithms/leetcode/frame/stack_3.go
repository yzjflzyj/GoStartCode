package main

/**
 * 496. 下一个更大元素 I
 * nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。
 * 给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。
 * 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，
 * 并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
 * 返回一个长度为nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
 */
//求nums2的单调栈,再建立nums1和nums2的对应关系
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indexMap := make(map[int]int)
	for i, v := range nums2 {
		indexMap[v] = i
	}
	greatNum := getNextGreaterElement(nums2)
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = greatNum[indexMap[v]]
	}
	return res
}

// 单调栈
func getNextGreaterElement(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	stack := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			res[i] = stack[len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return res
}

/**
  739.每日温度
  * 给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，
  * 其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。
*/
//单调栈,只是栈中记录的是索引
func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	res := make([]int, size)
	var stack []int
	for i := size - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}
		stack = append(stack, i)
	}
	return res
}

/**
 * 503. 下一个更大元素 II
 * 给定一个循环数组nums（nums[nums.length - 1]的下一个元素是nums[0]），返回nums中每个元素的 下一个更大元素 。
 * 数字 x的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
 * 如果不存在，则输出 -1。
 */
//处理环形数组:将数组重复一遍
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	var stack []int
	for i := 2*n - 1; i >= 0; i-- {
		val := nums[i%n]
		for len(stack) > 0 && stack[len(stack)-1] <= val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i%n] = stack[len(stack)-1]
		} else {
			res[i%n] = -1
		}
		stack = append(stack, val)
	}
	return res
}
func main() {
	dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70})
}
