package offer

import "github.com/Zelayan/go-fucking-written-examination/types"

func levelOrder(root *types.TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	// 记录每一层的数据
	lo := []*types.TreeNode{root}
	left := true
	// 记录的每一次的节点
	temp := new(types.TreeNode)
	for len(lo) > 0 {
		// 记录下一层的样子
		nextLayer := []*types.TreeNode{}
		// 记录当前一层的数据
		layerValue := []int{}
		for i := range lo {
			if left {
				temp = lo[i]
			} else {
				temp = lo[len(lo)-i-1]
			}
			if temp.Left != nil {

				nextLayer = append(nextLayer, temp.Left)
			}
			if temp.Right != nil {
				nextLayer = append(nextLayer, temp.Right)
			}
			layerValue = append(layerValue, lo[i].Val)
		}

		lo = nextLayer
		res = append(res, layerValue)
		left = !left
	}

	return res
}
