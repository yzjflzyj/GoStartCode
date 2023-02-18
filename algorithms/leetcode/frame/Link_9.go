package main

/**
 * 92. 反转链表 II
 * 给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre0 := dummyNode
	for i := 0; i < left-1; i++ {
		pre0 = pre0.Next
	}
	// cur0要反转的第一个，pre0要反转的前一个
	cur0 := pre0.Next
	pre := cur0
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// pre是反转的最后一个，cur是反转后面的第一个
	pre0.Next = pre
	cur0.Next = cur
	return dummyNode.Next
}

func main() {

}
