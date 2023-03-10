package main

/**
 * 26. 删除有序数组中的重复项
 */
//双指针
func removeDuplicates(nums []int) int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
