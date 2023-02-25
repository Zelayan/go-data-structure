package types

// TreeNode /**
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

const TreeNull = 10000

func GenListNode(arr []int) *ListNode {
	head := &ListNode{}
	curr := &ListNode{}
	curr.Next = head
	for _, v := range arr {
		curr.Next.Next = &ListNode{Val: v}
		curr.Next = curr.Next.Next
	}
	return head.Next
}

func BuildTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	// 每次都去处理队列第一个元素的孩子
	queue := []*TreeNode{root}

	// 这里的加2，其实就是处理队列第一个元素的孩子
	for i := 1; i < len(arr); i += 2 {
		node := queue[0]
		queue = queue[1:]

		if arr[i] != TreeNull {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		if i+1 < len(arr) && arr[i+1] != TreeNull {
			node.Right = &TreeNode{Val: arr[i+1]}
			queue = append(queue, node.Right)
		}
	}

	return root
}
