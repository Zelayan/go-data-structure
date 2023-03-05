package leetcode

type Trie struct {
	// 用数组存每个字符， 每个字符最多有26个还在
	children [26]*Trie
	// 是否是一个单词的结尾
	end bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		// 我们用0-26 来存储字母，-'a' 是让字母从0开始
		ch -= 'a'
		// 如果当前没有这个字母的话 就把对应下标的字母设置为当前字母
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		// node 节点进行切换
		node = node.children[ch]
	}
	// 当这个单词添加完毕，这时node处于最后一个字母，将end置为true
	node.end = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	// 如果找到最后，这个end字段为false, 说明不是单词的结尾
	return node != nil && node.end
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node != nil
}
