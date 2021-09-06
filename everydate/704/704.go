package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 5))
}

func search(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if i < len(nums) && nums[i] == target {
		return i
	} else {
		return -1

	}
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			low = mid + 1
		}
		if target < nums[mid] {
			high = mid - 1
		}
	}
	return -1
}
