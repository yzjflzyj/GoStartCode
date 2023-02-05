package main

import "container/heap"

/**
 * 23. 合并K个升序链表
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 */
//双指针
//k个有序节点,和21题类似，只是多个节点的比较，可以使用优先队列，用小顶堆实现
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := IntHeap{}
	heap.Init(&h)
	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}
	dummy := &ListNode{}
	p := dummy
	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		p.Next = min
		p = p.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return dummy.Next
}

type IntHeap []*ListNode

func (h IntHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Len() int            { return len(h) }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	ans := (*h)[n-1]
	*h = (*h)[:n-1]
	return ans
}

func main() {

}
