package main

/**
 * 5. 最长回文子串
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 */
//双指针:左右指针
func longestPalindrome(s string) string {
	res := ""
	for i := range s {
		str1 := getPalindrome(s, i, i)
		if len(str1) > len(res) {
			res = str1
		}
		str2 := getPalindrome(s, i, i+1)
		if len(str2) > len(res) {
			res = str2
		}
	}
	return res
}

func getPalindrome(s string, i, j int) string {
	for j < len(s) && i >= 0 && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}

func main() {

}
