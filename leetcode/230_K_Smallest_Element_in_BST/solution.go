package _30_K_Smallest_Element_in_BST

import "github.com/Zelayan/go-fucking-written-examination/types"

var (
	res  = 0
	rank = 0
)

func KsmElementInBST(root *types.TreeNode, k int) int {
	traverse(root, k)
	return res
}

func traverse(root *types.TreeNode, k int) {
	if root == nil {
		return
	}

	traverse(root.Left, k)
	rank++
	if rank == k {
		res = root.Val
		return
	}
	traverse(root.Right, k)

}
