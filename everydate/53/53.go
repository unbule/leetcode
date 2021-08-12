package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, -3}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	n := len(nums)
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
