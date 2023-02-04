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
func main() {

}
