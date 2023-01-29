package main

import "math"

/**
 * 104. 二叉树的最大深度
 */

// 方法1：dfs：前序遍历
func maxDepth(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cnt int) {
		if node == nil {
			return
		}
		cnt++
		ans = int(math.Max(float64(ans), float64(cnt)))
		dfs(node.Left, cnt)
		dfs(node.Right, cnt)
	}
	dfs(root, 0)
	return ans
}

// 方法2：动态规划
func maxDepth1(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth1(root.Left)
	rightDepth := maxDepth1(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth)) + 1)
}
func main() {

}
