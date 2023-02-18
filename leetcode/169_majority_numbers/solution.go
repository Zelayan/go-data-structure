package leetcode

// 利用hash表，遍历一次即可ss
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	resMap := make(map[int]int)
	n := len(nums)
	for i := range nums {
		if cur, ok := resMap[nums[i]]; ok {
			if (cur + 1) > n/2 {
				return nums[i]
			}
			resMap[nums[i]] = cur + 1
		} else {
			resMap[nums[i]] = 1
		}
	}
	return 0
}
