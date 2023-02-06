package main

/**
 * 19. 删除链表的倒数第 N 个结点
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	cur := pre
	first := pre
	for i := 0; i < n; i++ {
		pre = pre.Next
	}
	for pre.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return first.Next
}

func main() {

}
