// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
// 示例:
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // 返回 true
// trie.search("app");     // 返回 false
// trie.startsWith("app"); // 返回 true
// trie.insert("app");
// trie.search("app");     // 返回 true
// 说明:

// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。

package main

import "fmt"

func main() {
	fmt.Println([]byte("az"))
	// a := Trie{}
	// fmt.Println(a.Links)
}

// 链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/shi-xian-trie-qian-zhui-shu-by-leetcode/
type Trie struct {
	Links [26]*Trie
	isEnd bool
	Root  *Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := Trie{}
	return Trie{Root: &root}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.Root
	for i := 0; i < len(word); i++ {
		if node.Links[word[i]-97] == nil {
			node.Links[word[i]-97] = &Trie{}
		}
		node = node.Links[word[i]-97]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.Root
	for i := 0; i < len(word); i++ {
		if node.Links[word[i]-97] != nil {
			node = node.Links[word[i]-97]
		} else {
			return false
		}
	}
	// 树里包含这个word,但这个word之前有没有插入这个单词?
	// apple 里有app,但没有插入app之前app在树里不算单词
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for i := 0; i < len(prefix); i++ {
		if node.Links[prefix[i]-97] != nil {
			node = node.Links[prefix[i]-97]
		} else {
			return false
		}
	}
	// 比search简单,单词包含即可
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
