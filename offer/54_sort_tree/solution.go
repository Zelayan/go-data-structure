package leetcode

import (
	

	"github.com/Zelayan/go-fucking-written-examination/types"
)


var res []int

func kthLargest(root *types.TreeNode, k int) int {
    sort(root)
    return res[len(res)-k]
}

func sort(root *types.TreeNode)  {
    if root == nil {
        return
    }
    sort(root.Left)
    res =append(res, root.Val)
    sort(root.Right)
	
} 