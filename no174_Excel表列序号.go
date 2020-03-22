package main


//给定一个Excel表格中的列名称，返回其相应的列序号。
//例如，
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...

func titleToNumber(s string) int {
	ans := 0
	for _, c := range s {
		num := int(c-'A'+1)
		ans = ans*26+num
	}
	return ans
}
