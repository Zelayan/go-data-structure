package _05_Construct_Binary_Tree_from_Preorder_and_inorder_Traversal

import "github.com/Zelayan/go-fucking-written-examination/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *types.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := 0
	for i := range inorder {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	root := &types.TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:len(inorder[:index])+1])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
