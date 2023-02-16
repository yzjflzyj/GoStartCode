package main

/**
 * 142. 环形链表 II
 * 给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回null。
 * 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
 * 评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
 * 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
 * 不允许修改 链表。
 */
//链表，双指针，快慢指针判断成环,成环时，另一个指针从头开始同步遍历
func detectCycle(head *ListNode) *ListNode {
	//先快慢指针，相遇时快比慢指针多走了n圈，快指针走了2n圈，因此第二次再走a圈，相遇时，a即环开始的第一个节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 不成环
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 成环
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {

}
