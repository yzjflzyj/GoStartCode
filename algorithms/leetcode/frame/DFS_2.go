package main

/*
*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrack func(nums []int, start int)
	backtrack = func(path []int, start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, 0)
	return result
}

func main() {

}
