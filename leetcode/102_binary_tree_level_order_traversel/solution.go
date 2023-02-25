package leetcode

import (
	"github.com/Zelayan/go-fucking-written-examination/types"
)

func levelOrder(root *types.TreeNode) [][]int {

	var res [][]int
	if root == nil {
		return res
	}

	// 用于保存每一层
	queue := []*types.TreeNode{root}
	for {
		if len(queue) == 0 {
			break
		}
		var tempQueue []*types.TreeNode
		var tempRes []int
		for i := range queue {
			tempRes = append(tempRes, queue[i].Val)
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tempQueue = append(tempQueue, queue[i].Right)
			}
		}
		queue = tempQueue

		res = append(res, tempRes)

	}
	return res
}
