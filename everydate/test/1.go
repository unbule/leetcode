package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0}
	fmt.Println(wx(nums))
}

func wx(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for len(nums) > 1 {
		for i := 1; i < n; i += 2 {
			ans = ans + (nums[i] - nums[i-1])
			nums[i] = nums[i] ^ nums[i-1]
			nums[i-1] = 0
		}
		sort.Ints(nums)
		n = len(nums)
		nums = nums[n/2:]
	}
	return ans
}
