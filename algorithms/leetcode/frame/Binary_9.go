package main

/**
 * 889. 根据前序和后序遍历构造二叉树
 */

/*
利用前序第一个节点找根节点，利用前序的第2个节点去后序中比对确定根节点左子树的个数
预处理后序遍历生成node->index的映射
*/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postOrderIndex := make(map[int]int, len(postorder))
	for k, v := range postorder {
		postOrderIndex[v] = k
	}
	return dfsBuildPrePost(0, 0, len(preorder), preorder, postOrderIndex)
}

/*
*
@param preStart  前序遍历的开始索引
@param postStart 后序遍历开始索引
@param N 树中节点个数
@param preOrder []int 后序遍历节点副本
@param postOrderIndex []int 后序遍历节点到offset的映射
*/
func dfsBuildPrePost(preStart int, postStart int, N int, preOrder []int, postOrderIndex map[int]int) *TreeNode {
	if N == 0 {
		return nil
	}
	root := &TreeNode{Val: preOrder[preStart]}
	if N == 1 {
		return root
	}
	leftNode := preOrder[preStart+1]
	L := postOrderIndex[leftNode] - postStart + 1
	root.Left = dfsBuildPrePost(preStart+1, postStart, L, preOrder, postOrderIndex)
	root.Right = dfsBuildPrePost(preStart+L+1, postStart+L, N-1-L, preOrder, postOrderIndex)
	return root
}

func main() {

}
