package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

func levelOrder(root *types.Node) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int

	queue := []*types.Node{root}

	for {
		if len(queue) == 0 {
			break
		}
		var tempQueue []*types.Node
		tempRes := []int{}
		for i := range queue {
			tempRes = append(tempRes, queue[i].Val)
			tempQueue = append(tempQueue, queue[i].Children...)
		}
		queue = tempQueue
		res = append(res, tempRes)

	}
	return res

}
