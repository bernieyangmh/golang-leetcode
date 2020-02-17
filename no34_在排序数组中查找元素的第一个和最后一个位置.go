package main

import "fmt"

func main() {
	a := []int{2,2}
	fmt.Println(binarySearch(a, 0, 1, 3))

}

func binarySearch(nums []int, s, e, target int) int {
	if s > e {
		return -1
	}
	mid := (s + e) / 2

	if nums[mid] < target {
		return binarySearch(nums, mid+1, e, target)
	} else if nums[mid] > target{
		return binarySearch(nums, s, mid-1, target)
	} else {
		return mid
	}
}


func binarySearch2(sortedArray []int, lookingFor int) int {
    var low int = 0
    var high int = len(sortedArray) - 1
    for low <= high {
        var mid int =low + (high - low)/2
        var midValue int = sortedArray[mid]
        if midValue == lookingFor {
            return mid
        } else if midValue > lookingFor {
            high = mid -1
        } else {
            low = mid + 1
        }
    }
    return -1
}