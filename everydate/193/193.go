package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	f, s := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		f, s = s, max(f+v, s)
	}
	return s
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
