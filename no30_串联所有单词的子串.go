package main

//
//给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
//示例 1：
//输入：
//s = "barfoothefoobarman",
//words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
//示例 2：
//输入：
//s = "wordgoodgoodgoodbestword",
//words = ["word","good","best","word"]
//输出：[]

func findSubstring(s string, words []string) []int {
	res := []int{}

	wordNum := len(words)
	if wordNum == 0 {
		return res
	}
	wordlen := len(words[0])
	allWords := make(map[string]int)

	// map保存所有单词出现的次数
	for _, w := range words {
		if _, ok := allWords[w]; ok {
			allWords[w] += 1
		} else {
			allWords[w] = 1
		}
	}

	// 从头开始遍历
	for i := 0; i < len(s)-wordNum*wordlen+1; i++ {
		hasWords := make(map[string]int)
		num := 0
		// 如果我当前找到的单词数量等于目标单词数量，就可以跳出来判断了
		for num < wordNum {
			// 每次取单词的方法，
			word := s[i+num*wordlen : i+(num+1)*wordlen]
			// 如果取到的单词在目标单词里面，看下数量有没有超
			// 如果不在，字符串不符合，直接break
			if _, ok := allWords[word]; ok {
				var value int
				if _, ok := hasWords[word]; ok {
					value = hasWords[word]
				} else {
					value = 0
				}
				hasWords[word] = value + 1
				if hasWords[word] > allWords[word] {
					break
				}
			} else {
				break
			}
			num++
		}
		// 当数量一致，把i放进结果
		if num == wordNum {
			res = append(res, i)
		}
	}
	return res
}

func findSubstring2(s string, words []string) []int {
	res := []int{}

	wordNum := len(words)
	if wordNum == 0 {
		return res
	}
	wordlen := len(words[0])
	allWords := make(map[string]int)

	// map保存所有单词出现的次数
	for _, w := range words {
		if _, ok := allWords[w]; ok {
			allWords[w] += 1
		} else {
			allWords[w] = 1
		}
	}
	// 单词有多长，起始的索引就有多少种
	for j := 0; j < wordlen; j++ {

		hasWords := make(map[string]int)
		num := 0

		//从开始走到最后能满足所有单词长度的索引，每次递增一个单词的长度
		for i := j; i < len(s)-wordNum*wordlen+1; i = i + wordlen {
			hasRemoved := false

			for num < wordNum {
				word := s[i+num*wordlen : i+(num+1)*wordlen]
				if _, ok := allWords[word]; ok {
					var value int
					if _, ok := hasWords[word]; ok {
						value = hasWords[word]
					} else {
						value = 0
					}
					hasWords[word] = value + 1
					//3.如果单词符合，但是次数超了
					if hasWords[word] > allWords[word] {
						hasRemoved = true
						removeNum := 0
						// 一直移除单词，直到次数符合，相当于i不停向右走
						for hasWords[word] > allWords[word] {
							firstword := s[i+removeNum*wordlen : i+(removeNum+1)*wordlen]
							hasWords[firstword] = hasWords[firstword] - 1
							removeNum++
						}
						// 移除了removeNum个单词，新增了一个word
						num = num - removeNum + 1
						// i应该移动removeNum个长度，比如去掉两个单词，就移动6位。
						// 但因为最外层for循环i本身就移动一个单词长度，所以减1
						i = i + (removeNum-1)*wordlen
						break
					}
					// 单词不匹配，直接从当前单词开始，最外层的for循环会移动到下一个单词
				} else {
					hasWords = make(map[string]int)
					i = i + num*wordlen
					num = 0
					break
				}
				num++
			}
			if num == wordNum {
				res = append(res, i)
			}
			//如果完全匹配，去掉第一个单词
			if num > 0 && !hasRemoved {
				firstWord := s[i : i+wordlen]
				hasWords[firstWord] = hasWords[firstWord] - 1
				num--
			}
		}

	}
	return res
}
