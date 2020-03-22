package main


func findRepeatedDnaSequences(s string) []string {
	length := len(s)
	res := []string{}
	set := make(map[string]bool)
	for i:=0;i<=length-10;i++{
		key := s[i:i+10]
		if _, ok := set[key]; ok {
			res = append(res, key)
		} else {
			set[key] = true
		}
	}
	return res
}
