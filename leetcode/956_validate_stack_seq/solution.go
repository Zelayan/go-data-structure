package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 2 {
		return true
	}
	h := popped[0]
	var stack []int
	check := -1
	for i := range pushed {
		if pushed[i] != h {
			stack = append(stack, pushed[i])
		} else {
			check = i
			break
		}
	}
	for i := 1; i < len(popped); i++ {
		if len(stack) != 0 && popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			find := false
			for j := check + 1; j < len(popped); j++ {
				if pushed[j] != popped[i] {
					stack = append(stack, pushed[j])
					check = check + 1
				} else {
					check = check + 1
					find = true
					break
				}
			}
			if !find {
				return find
			}
		}
	}
	return true
}

// 使用模拟栈的方式去实现
func validateStackSequences2(pushed []int, popped []int) bool {
	if len(pushed) == 2 {
		return true
	}
	var stack []int
	j := 0
	for i := range pushed {
		// 入栈
		stack = append(stack, pushed[i])
		// 判断栈顶元素是否相等
		for len(stack) != 0 && stack[len(stack)-1] == popped[j] {
			// 如果相等就出栈
			stack = stack[:len(stack)-1]
			j++
		}
	}
	// 如果栈里面还有元素说明有问题
	return len(stack) == 0
}
