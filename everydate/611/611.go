package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
}

func triangleNumber(nums []int) int {
	n := len(nums)
	ans := 0
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			ans += sort.SearchInts(nums[j+1:], nums[j]+nums[i])
		}
	}
	return ans
}
