package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	//n := len(nums)
	//res := 0
	for i := 1; i < len(nums); i++ {
		//res = res + nums[i]
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
