// 设计一个支持以下两种操作的数据结构：
// void addWord(word)
// bool search(word)
// search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。
// 示例:
// addWord("bad")
// addWord("dad")
// addWord("mad")
// search("pad") -> false
// search("bad") -> true
// search(".ad") -> true
// search("b..") -> true
// 说明:
// 你可以假设所有单词都是由小写字母 a-z 组成的。

package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Println(obj.Search("bad"))
}

type WordDictionary struct {
	Links [26]*WordDictionary
	isEnd bool
	Root  *WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	root := WordDictionary{}
	return WordDictionary{Root: &root}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.Root
	for i := 0; i < len(word); i++ {
		if node.Links[word[i]-97] == nil {
			node.Links[word[i]-97] = &WordDictionary{}
		}
		node = node.Links[word[i]-97]
	}
	node.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	// 开始传的是根根结点
	return searchHelper(word, this.Root)
}

func searchHelper(word string, node *WordDictionary) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == 46 {
			for _, n := range node.Links {
				if n == nil {
					continue
				}
				if searchHelper(word[i+1:], n) {
					return true
				}
			}
			return false
		}
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

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
