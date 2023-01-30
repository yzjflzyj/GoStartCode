package main

/**
 * 226. 翻转二叉树
 *给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 */
//剑指27的二叉树的镜像,相同
//动态规划型,DFS

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

func main() {

}
