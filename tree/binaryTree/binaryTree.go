package binarytree

//定义二叉树结构体
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

//创建二叉树
func CreateBinaryTree(i int, nums []int) *BinaryTree {
	tree := &BinaryTree{nums[i], nil, nil}
	//左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = CreateBinaryTree(2*i+1, nums)
	}
	//右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Right = CreateBinaryTree(2*i+2, nums)
	}
	return tree
}
