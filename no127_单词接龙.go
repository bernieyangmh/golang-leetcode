package main

//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。
//转换需遵循如下规则：
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典中的单词。
//说明:
//如果不存在这样的转换序列，返回 0。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

type node struct {
	word string
	state int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	l := len(beginWord)

	dict := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < l; j++ {
			k := wordList[i][:j] + "*" + wordList[i][j+1:]
			dict[k] = append(dict[k], wordList[i])
		}
	}

	queue := []node{node{word: beginWord, state: 1}}
	set := make(map[string]struct{})
	set[beginWord] = struct{}{}

	for len(queue) != 0 {
		current_node := queue[0]
		queue = queue[1:]
		for i := 0; i < l; i++ {
			intermediate_word := current_node.word[:i] + "*" + current_node.word[i+1:]
			for j := 0; j < len(dict[intermediate_word]); j++ {
				if dict[intermediate_word][j] == endWord {
					return current_node.state + 1
				}
				if  _, ok := set[dict[intermediate_word][j]]; !ok {
					set[dict[intermediate_word][j]] = struct{}{}
					queue = append(queue, node{word: dict[intermediate_word][j], state: current_node.state + 1})
				}
			}
		}
	}
	return 0
}