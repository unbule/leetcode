package main

import "fmt"

func main() {
	nums := []int{7, 7, 7, 7, 7}
	fmt.Println(numberOfArithmeticSlices(nums))
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	a, res := 0, 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			a = a + 1
		} else {
			a = 0
		}
		res = res + a
	}
	return res
}
