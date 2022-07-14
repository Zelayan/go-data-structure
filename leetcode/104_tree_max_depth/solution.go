package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	depth = 0
	// 用来保存最大的深度
	res = 0
)

func maxDepth(root *TreeNode) int {

	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root != nil {
		return
	}

	// 进入一个节点加一
	depth++
	res = max(res, depth)
	traverse(root.Left)
	traverse(root.Left)

	// 出去一个节点减1
	depth--
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// TreeNode /**
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
