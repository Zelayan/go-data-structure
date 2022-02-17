package linknodes

type LinkNode struct {
	Data int64
	Next *LinkNode
}

func (l *LinkNode) Add(next int64) {
	ln := new(LinkNode)
	l.Next = 
}