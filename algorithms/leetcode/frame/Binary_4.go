package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * 116. 填充每个节点的下一个右侧节点指针
 */
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var dfsFunc func(*Node, *Node)
	dfsFunc = func(node1 *Node, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}
		node1.Next = node2
		dfsFunc(node1.Left, node1.Right)
		dfsFunc(node1.Right, node2.Left)
		dfsFunc(node2.Left, node2.Right)
	}
	dfsFunc(root.Left, root.Right)
	return root
}

func main() {

}
