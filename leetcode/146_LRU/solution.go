package leetcode

type node struct {
	key int
	value int
	Next *node
	Prev *node
}

type LRUCache struct {
	cap int
	length int
	LinkedList *node
	hmap map[int]*node
}


func Constructor(capacity int) LRUCache {
	linkedHead := &node{}
	linkedHead.Next = linkedHead
	linkedHead.Prev = linkedHead
	return LRUCache{
		cap: capacity,
		LinkedList: linkedHead,
		hmap: make(map[int]*node, capacity),
	}
}

func (this *LRUCache) MoveHead(target *node) {
	this.DelNode(target)
	this.AddNodeToHead(target)
}

func (this *LRUCache) DelNode(target *node) {
	target.Prev.Next = target.Next
	target.Next.Prev = target.Prev
}

func (this *LRUCache) AddNodeToHead(target *node) {
	this.LinkedList.Next.Prev = target
	target.Next = this.LinkedList.Next
	target.Prev = this.LinkedList
	this.LinkedList.Next = target
}

func (this *LRUCache) Get(key int) int {
	target, ok := this.hmap[key]
	if !ok {
		return -1
	}
	this.MoveHead(target)
	return target.value
}


func (this *LRUCache) Put(key int, value int)  {
	target, ok := this.hmap[key]
	if ok {
		target.value = value
		this.MoveHead(target)
		return
	}

	if this.length == this.cap {

	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */