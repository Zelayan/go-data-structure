package leetcode

// 左右指针分别从左右最有开始，如果相加大于sum，则需要将r-- 如果小于sum 则r++
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	nums := 0
	for l < r {

		nums = numbers[l] + numbers[r]
		if nums == target {
			return []int{l + 1, r + 1}
		} else if nums < target {
			l++
		} else {
			r--
		}
	}

	return []int{-1, -1}
}
