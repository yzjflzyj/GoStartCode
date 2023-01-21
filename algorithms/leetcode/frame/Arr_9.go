package main

/**
 * 283. 移动零
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 */
//双指针
func moveZeroes(nums []int) {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[res] = nums[i]
			res++
		}
	}
	for i := res; i < length; i++ {
		nums[i] = 0
	}
}

func main() {

}
