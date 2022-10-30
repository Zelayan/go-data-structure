package leetcode

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *types.TreeNode) int {
    // 一个节点的直径就是 它的左子树加右子树高度
    var res int
    var dfs func(*types.TreeNode) int 
	dfs = func(node *types.TreeNode) int {
		if node == nil {
			return 0
		}
        // 计算左子树的高度
        left := dfs(node.Left)
        // 计算右子树的高度
        right := dfs(node.Right)
        // 如果左子树 + 右子树的高度大于全局的高度，就刷新
        if left + right > res {
            res = left +right
        }
        if left > right {
            return left + 1
        }
        return right + 1
        
    }

    dfs(root)
    return res
}