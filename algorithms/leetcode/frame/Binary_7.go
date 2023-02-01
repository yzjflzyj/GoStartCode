package main

/**
 * 105. 从前序与中序遍历序列构造二叉树
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 复用中序的位置记录
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}
	var build func([]int, []int, int, int, int, int) *TreeNode
	build = func(pre, in []int, preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft > preRight {
			return nil
		}
		leftSize := indexMap[pre[preLeft]] - inLeft
		leftNode := build(pre, in, preLeft+1, preLeft+leftSize, inLeft, inLeft+leftSize-1)
		rightNode := build(pre, in, preLeft+leftSize+1, preRight, indexMap[pre[preLeft]]+1, inRight)
		return &TreeNode{pre[preLeft], leftNode, rightNode}
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	// 左子树的长度left
	left := findRootIndex(preorder[0], inorder)
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree1(preorder[1:left+1], inorder[:left]),
		Right: buildTree1(preorder[left+1:], inorder[left+1:])}
	return root
}
func findRootIndex(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

func main() {

}
