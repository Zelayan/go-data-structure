package linknodes

import "fmt"

// single linkNode
type LinkNode struct {
	Data int64
	NextNode *LinkNode
}

func (l *LinkNode) Add(next int64) {
	nextNode := new(LinkNode)
	nextNode.Data = next
	l.NextNode = nextNode
}

func (l *LinkNode) Show() {
	
	nowNode := l
	for {
		if nowNode != nil {
			fmt.Print(nowNode.Data)
			nowNode = nowNode.NextNode
			continue
		}
		break
	}
}

// 
type Ring struct {
	next, pre *Ring
	Value interface{}
}

func (r *Ring) Init() *Ring{
	r.next = r
	r.pre = r
	r.Value = 
}