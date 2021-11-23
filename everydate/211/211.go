package main

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (t *TrieNode) insert(word string) {
	node := t
	for _, ch := range word {
		ch = ch - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

type WordDictionary struct {
	trieroot *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.trieroot.insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, tn *TrieNode) bool {
		if index == len(word) {
			return tn.isEnd
		}
		ch := word[index]
		if ch != '.' {
			child := tn.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range tn.children {
				child := tn.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, this.trieroot)
}
