package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	res := 0
	start := 0
	for i, v := range s {
		// 如果之前的字串中有，则将start 设置为之前的后一个
		if pre, ok := m[v]; ok {
			start = max(start, pre+1)
		}
		// 结果就是左边-右边+1 和之前的答案里面最大的
		res = max(res, i-start+1)
		m[v] = i
	}
	return res
}
