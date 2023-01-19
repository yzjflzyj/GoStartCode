package main

// 83. 删除排序链表中的重复元素
// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
// 双指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
