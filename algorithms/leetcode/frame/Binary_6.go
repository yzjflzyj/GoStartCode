package main

/**
 * 654. 最大二叉树
 * 给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:
 * 创建一个根节点，其值为nums 中的最大值。
 * 递归地在最大值左边的子数组前缀上构建左子树。
 * 递归地在最大值 右边 的子数组后缀上构建右子树。
 * 返回nums 构建的 最大二叉树 。
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var f func(start, end int) *TreeNode
	f = func(start, end int) *TreeNode {
		if start >= end {
			return nil
		} // [start, end) 左闭右开区间
		max := start
		for i := start; i < end; i++ {
			if nums[i] > nums[max] {
				max = i
			}
		}
		return &TreeNode{nums[max], f(start, max), f(max+1, end)}
	}
	return f(0, len(nums))
}
func main() {

}
