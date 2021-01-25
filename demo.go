package main

func main() {


}

func upper_bound_( n int ,  v int ,  a []int ) int {
	l, r := 0, n
	for l < r {
		mid := l +(r-l)/2
		if a[mid] < v {
			l = mid+1
		}else {
			r = mid
		}
	}
	return r+1
}