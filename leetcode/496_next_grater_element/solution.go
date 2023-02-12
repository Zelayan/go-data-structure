package leetcode

//
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := nextGreaterElement2(nums2)

	var res []int
	for i := range nums1 {
		if ans, ok := m[nums1[i]]; ok {
			res = append(res, ans)
		}
	}
	return res
}

// 在数组中找后续比当前大的第一个数
func nextGreaterElement2(num []int) map[int]int {
	res := make(map[int]int, len(num))
	var stack []int
	for i, v := range num {
		// 先默认都没有找到
		res[num[i]] = -1
		// 如果栈为空，先入栈
		if len(stack) == 0 {
			stack = append(stack, num[i])
		} else {
			for {
				// 如果栈中有数据，判断当前要判断的数和栈顶数据的关系，如果当前数 大于栈顶数据，说明当前数是栈顶数据的第一个大数
				// 将栈顶元素抛出，并将结果记录
				if len(stack) != 0 && stack[len(stack)-1] < v {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					res[top] = num[i]
				} else {
					// 如果栈为空，则将当前数据压入栈中
					stack = append(stack, num[i])
					break
				}
			}
		}
	}
	return res
}
