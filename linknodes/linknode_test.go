package linknodes

import "testing"

func TestShow(t *testing.T) {
	
}	
func ExampleShow() {
	last := new(LinkNode)
	last.Data = 30

	second := new(LinkNode)
	second.Data = 20
	second.NextNode = last
	head := new(LinkNode)
	head.Data=10
	head.NextNode = second

	head.Show()
	//Output:102030
	
}

