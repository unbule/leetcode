package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement2(nums))
}

func findPeakElement(nums []int) int {
	n := len(nums)
	i := 1
	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	for i = 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			break
		}
	}
	return i
}

func findPeakElement2(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i >= n || i == -1 {
			return math.MinInt64
		}
		return nums[i]
	}
	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		} else if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
