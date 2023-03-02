package leetcode

func longestPalindrome(s string) string {

	if s == "" {
		return ""
	}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// 中间字母是一个
		left1, right1 := palindrom(s, i, i)
		// 中间字母是两个想同的
		left2, right2 := palindrom(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func palindrom(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}

	return left + 1, right - 1
}
