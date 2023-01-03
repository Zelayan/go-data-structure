package leetcode

func singleNumber(nums []int) int {
	single := 0
	for _, i := range nums {
		single ^= i
	}
	return single
}
