package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	d := make(map[float64]float64)
	w := float64((t)) + 1
	i := 0
	for i < len(nums) {
		i++
		m := getId(float64(nums[i]), w)
		if _, ok := d[m]; ok {
			return true
		}
		if m1, ok := d[m-1]; ok && abs(float64(nums[i])-m1) < w {
			return true
		}
		if m2, ok := d[m+1]; ok && abs(float64(nums[i])-m2) < w {
			return true
		}
		d[m] = float64(nums[i])
		if i >= k {
			deleteKey := getId(float64(nums[i-k]), w)
			delete(d, deleteKey)
		}
	}
	return false
}

func getId(x, w float64) float64 {
	if x < 0 {
		return (x + 1) / (w - 1)
	}
	return x / w
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x

}
