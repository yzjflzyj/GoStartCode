package main

/**
 * 111. 二叉树的最小深度
 */
//BFS算法框架
// 计算从起点 start 到终点 target 的最近距离
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//切片实现队列
	q := []*TreeNode{root}
	depth := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth
}

func main() {

}
