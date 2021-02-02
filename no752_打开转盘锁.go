package main

func openLock(deadends []string, target string) int {
	Visited := make(map[string]struct{})
	for _, d := range deadends {
		Visited[d] = struct{}{}
	}
	q := []string{}
	step := 0
	q = append(q, "0000")
	Visited["0000"] = struct{}{}

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			cur := q[size-1]
			q := q[:size-1]
			if _, ok := Visited[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			//对4位分别转上下两次
			for j := 0; j < 4; j++ {
				up, down := turnOnce(cur, j)
				if _, ok := Visited[up]; !ok {
					q = append(q, up)
					Visited[up] = struct{}{}
				}
				if _, ok := Visited[down]; !ok {
					q = append(q, down)
					Visited[down] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

func turnOnce(s string, bit int) (string, string) {
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var s1, s2 string
	n := s[bit] - '0'
	s1 = nums[(n+1)%10]
	s2 = nums[(n+9)%10]
	switch bit {
	case 0:
		return s1 + s[1:], s2 + s[1:]
	case 3:
		return s[0:3] + s1, s[0:3] + s2
	default:
		return s[0:bit] + s1 + s[bit+1:], s[0:bit] + s2 + s[bit+1:]
	}
}
