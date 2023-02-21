package leetcode

func findDisappearedNumbers(nums []int) []int {

	var res []int
	base := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		base[i+1] = i + 1
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := base[nums[i]]; ok {
			base[nums[i]] = 0
		}
	}

	for i := 0; i < len(nums); i++ {
		if base[i+1] != 0 {
			res = append(res, i+1)
		}
	}

	return res
}
