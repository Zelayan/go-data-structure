package leetcode

func trap(nums []int) int {

	if len(nums) == 1 || len(nums) == 2 {
		return 0
	}
	var stack []int
	max := -1
	res := 0

	for i := 0; i < len(nums); i++ {
		// 当栈里面没有元素时， 如果元素不为0，直接入栈
		if len(stack) == 0 {
			if nums[i] != 0 {
				stack = append(stack, nums[i])
				max = nums[i]
			}
		} else if len(stack) == 1 { // 如果栈里面有一个元素，如果当前元素大于 栈顶元素，直接进行替换，
			if nums[i] > stack[len(stack)-1] {
				stack = []int{nums[i]}
				max = nums[i]
			} else {
				// 如果小于，直接入栈
				stack = append(stack, nums[i])
			}
		} else {
			// 如果当前元素大于栈顶元素
			if nums[i] > stack[len(stack)-1] {
				// 如果当前元素大于栈低元素，需要进行计算面积
				if nums[i] >= max {
					stack = append(stack, nums[i])
					res = res + getArea(stack)
					stack = []int{nums[i]}
					max = nums[i]
				} else {
					// 负责进行入展
					stack = append(stack, nums[i])
				}
			} else {
				stack = append(stack, nums[i])
			}
		}
	}

	if len(stack) >= 3 {
		res = res + getArea(stack)
	}
	return res
}

func getArea(nums []int) int {

	max := nums[len(nums)-1]
	right := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= max {
			right = i
		}
	}

	length := len(nums) - 2
	height := nums[0]
	if nums[0] > nums[len(nums)-1] {
		height = nums[len(nums)-1]
	}
	area := height * length
	for i := 1; i < len(nums)-1; i++ {
		area = area - nums[i]
	}
	if area < 0 {
		area = 0
	}
	return area
}
