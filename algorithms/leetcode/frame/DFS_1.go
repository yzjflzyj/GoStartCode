package main

/**
 * 46.全排列
 */
//DFS的交换的回溯算法,见Test_46,交换的方式,效率更高
//剑指38,去重的全排列,也是DFS交换的方式
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	path := make([]int, len(nums))
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(path, used)
	return res
}

func main() {

}
