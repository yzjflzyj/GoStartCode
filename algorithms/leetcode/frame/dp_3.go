package main

import "sort"

// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {
	//使得f按照升序排列
	f := []int{}
	for _, e := range nums {
		//h := e
		//golang内置的SearchInts函数已经实现了二分查找
		//使用golang内置的SearchInts函数来查找当前队列中是否有，如果没有，返回此数应该被插入的位置。
		if i := sort.SearchInts(f, e); i < len(f) {
			f[i] = e
		} else {
			f = append(f, e)
		}
	}
	return len(f)
}
