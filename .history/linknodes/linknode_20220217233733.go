package linknodes

import "fmt"

type LinkNode struct {
	Data int64
	Next *LinkNode
}

func (l *LinkNode) Add(next int64) {
	nextNode := new(LinkNode)
	nextNode.Data = next
	l.Next = nextNode
}

func (l *LinkNode) Show() {
	
	fmt.Print(l.Data)
	nowNode := l
	for {
		if nowNode != nil {
			fmt.Print(nowNode.Data)
			nowNode = nowNode.Next
			continue
		}
		break
	}
}