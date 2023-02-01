package main

/**
 *295. 数据流的中位数
 */

type MedianFinder struct {
	lowHeap  []int //偏小一边的大顶堆
	highHeap []int //偏大一边的小顶堆
}

func Constructor_1() MedianFinder {
	finder := MedianFinder{
		lowHeap:  make([]int, 0, 5),
		highHeap: make([]int, 0, 5),
	}
	return finder
}

// 两个堆的调整过程与平衡二叉树平衡过程类型
// 同样是通过旋转保持平衡 只是操作更为简单
func (finder *MedianFinder) AddNum(num int) {
	var length1, length2 = len(finder.lowHeap), len(finder.highHeap)
	var top1, top2 = 0, 0
	if length1 > 0 {
		top1 = finder.lowHeap[0]
	}
	if length2 > 0 {
		top2 = finder.highHeap[0]
	}

	if length1 == length2 {
		// 两堆大小相等时
		// num大于偏小一边的堆最大值,则插入偏大一边
		// 反之插入偏大一边
		if num >= top1 {
			finder.highHeap = append(finder.highHeap, num)
			AdjustBotToTopSmall(finder.highHeap)
		} else if num < top2 {
			finder.lowHeap = append(finder.lowHeap, num)
			AdjustBotToTopLarge(finder.lowHeap)
		}
	} else if length1 > length2 {
		// 两堆大小不相等时
		// 必须插入堆元素更小的一边 保持中位数的平衡
		if num >= top1 {
			finder.highHeap = append(finder.highHeap, num)
			AdjustBotToTopSmall(finder.highHeap)
		} else {
			// 如果num恰好位于另一边 进行两次调整操作
			// 旋转一边的堆顶插入另一半 将num插入这一边
			finder.highHeap = append(finder.highHeap, top1)
			AdjustBotToTopSmall(finder.highHeap)
			finder.lowHeap[0] = num
			AdjustTopToBotLarge(finder.lowHeap)
		}
	} else {
		// 与情况二对称
		if num <= top2 {
			finder.lowHeap = append(finder.lowHeap, num)
			AdjustBotToTopLarge(finder.lowHeap)
		} else {
			finder.lowHeap = append(finder.lowHeap, top2)
			AdjustBotToTopLarge(finder.lowHeap)
			finder.highHeap[0] = num
			AdjustTopToBotSmall(finder.highHeap)
		}
	}
}

func (finder *MedianFinder) FindMedian() float64 {
	var length1, length2 = len(finder.lowHeap), len(finder.highHeap)
	var result float64
	//保证两堆元素数量相等 由两堆顶显然可推中位数
	if length1 > length2 {
		result = float64(finder.lowHeap[0])
	} else if length1 < length2 {
		result = float64(finder.highHeap[0])
	} else {
		result = (float64(finder.lowHeap[0]) + float64(finder.highHeap[0])) / 2
	}
	return result
}

// 大顶堆 自上而下调整算法
func AdjustTopToBotLarge(heap []int) {
	var topValue = heap[0]
	var top = 0
	for i := 1; i < len(heap); i = i*2 + 1 {
		if i < len(heap)-1 && heap[i+1] > heap[i] {
			i++
		}
		if topValue < heap[i] {
			heap[top] = heap[i]
			top = i
		} else {
			break
		}
	}
	heap[top] = topValue
}

// 大顶堆 自下而上调整算法
func AdjustBotToTopLarge(heap []int) {
	var botValue = heap[len(heap)-1]
	var bot = len(heap) - 1
	for i := (bot - 1) / 2; i >= 0; i = (i - 1) / 2 {
		if botValue > heap[i] {
			heap[bot] = heap[i]
			bot = i
		} else {
			break
		}
		if i == 0 {
			break
		}
	}
	heap[bot] = botValue
}

// 小顶堆 自上而下调整算法
func AdjustTopToBotSmall(heap []int) {
	var topValue = heap[0]
	var top = 0
	for i := 1; i < len(heap); i = i*2 + 1 {
		if i < len(heap)-1 && heap[i+1] < heap[i] {
			i++
		}
		if topValue > heap[i] {
			heap[top] = heap[i]
			top = i
		} else {
			break
		}
	}
	heap[top] = topValue
}

// 小顶堆 自下而上调整算法
func AdjustBotToTopSmall(heap []int) {
	var botValue = heap[len(heap)-1]
	var bot = len(heap) - 1
	for i := (bot - 1) / 2; i >= 0; i = (i - 1) / 2 {
		if botValue < heap[i] {
			heap[bot] = heap[i]
			bot = i
		} else {
			break
		}
		if i == 0 {
			break
		}
	}
	heap[bot] = botValue
}
