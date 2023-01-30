package main

import "math"

/**
 *543. 二叉树的直径
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 */
//⼀旦你发现题⽬和⼦树有关，那⼤概率要给函数设置合理的定义和返回值，在后序位置写代码了。即动态规划二叉树
// 记录最⼤直径的⻓度
func diameterOfBinaryTree(root *TreeNode) (res int) {
	var getMaxDepth func(*TreeNode) int
	// 方法定义：以当前节点为根结点，得到的最长路径
	getMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := getMaxDepth(node.Left)
		rightMax := getMaxDepth(node.Right)
		// 收集结果
		res = int(math.Max(float64(res), float64(leftMax+rightMax)))

		max := math.Max(float64(leftMax), float64(rightMax))
		return int(max + 1)
	}
	getMaxDepth(root)
	return res
}

func main() {

}
