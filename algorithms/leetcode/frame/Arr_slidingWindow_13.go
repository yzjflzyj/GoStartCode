package main

import "math"

/**
 * 76. 最小覆盖子串
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
 * 注意：
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 */

//滑动窗口模板型方法，主要处理子串类问题：固定流程
/**
 *1.右侧扩大窗口
 *2.进⾏窗⼝内数据的⼀系列更新
 *3.判断左侧窗⼝是否要收缩
 *4.进⾏窗⼝内数据的⼀系列更新
 */
func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		//1.扩大右边界
		c := s[right]
		right++
		//2.更新窗口内数据
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//3.判断是否需要收缩左边界
		for len(need) == valid {
			if length > right-left {
				length = right - left
				start = left
			}
			//4.左侧移出,并对应更新数据
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
