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
