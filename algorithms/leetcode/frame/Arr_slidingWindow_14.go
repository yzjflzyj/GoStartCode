package main

/**
 * 567. 字符串的排列
 * 给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 */

//滑动窗口模板型方法，主要处理子串类问题：固定流程

/**
 * 1.右侧扩大窗口
 * 2.进⾏窗⼝内数据的⼀系列更新
 * 3.判断左侧窗⼝是否要收缩
 * 4.进⾏窗⼝内数据的⼀系列更新
 */
// 判断 s 中是否存在 t 的排列
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		//1.扩大右边界
		c := s2[right]
		right++
		//2.更新窗口内数据
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//3.判断是否需要收缩左边界
		for right-left >= len(s1) {
			// 在这⾥判断是否找到了合法的⼦串
			if valid == len(need) {
				return true
			}
			//4.左侧移出,并对应更新数据
			d := s2[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

func main() {

}
