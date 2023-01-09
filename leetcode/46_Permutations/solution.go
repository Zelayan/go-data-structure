package leetcode

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums, len(nums), []int{})
	return res
}

// backTrack
//
func backTrack(nums []int, numsLen int, path []int) {
	// 说明子树的一条路径遍历完毕
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		cur := nums[i]
		path = append(path, cur)
		// 做选择
		// 将当前的使用的摘取， 例如这一步选择了1 就1 摘取，只留2和3 去遍历子树
		nums = append(nums[:i], nums[i+1:]...)
		backTrack(nums, len(nums), path)

		// 回溯 将摘取的数据归还
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]
	}
}
