package leetcode

import (
	"github.com/Zelayan/go-fucking-written-examination/types"
)

func findTilt(root *types.TreeNode) int {
    if root == nil {
        return 0
    }
	ans := 0
	var dfs func(*types.TreeNode) int
	dfs = func(node *types.TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		// 如果左右孩子没有，则坡度为0
		// 这里记录总的坡度
		ans += abs(sumLeft - sumRight)
		// 返回此节点的值，包括左右孩子和自己
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    } 
    return x
}