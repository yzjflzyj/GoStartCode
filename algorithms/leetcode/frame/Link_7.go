package main

/**
 * 160. 相交链表
 * 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
 */
//与剑指52相同
//链表分别为a+c,b+c.两个链表跑完a+c+b和b+a+c,终点则是相同的c的头节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = headB
		} else {
			cur1 = cur1.Next
		}
		if cur2 == nil {
			cur2 = headA
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}

func main() {

}
