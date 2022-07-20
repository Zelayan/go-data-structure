package _54_maximum_binary_tree

import "github.com/Zelayan/go-fucking-written-examination/types"

func constructMaximumBinaryTree(nums []int) *types.TreeNode {
	if nums == nil {
		return nil
	}
	if len(nums) == 0 {
		return nil
	}

	max, i := max(nums)

	result := &types.TreeNode{}
	result.Val = max

	result.Left = constructMaximumBinaryTree(nums[:i])
	result.Right = constructMaximumBinaryTree(nums[i+1:])

	return result
}

func max(nums []int) (int, int) {
	max := nums[0]
	index := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return max, index
}
