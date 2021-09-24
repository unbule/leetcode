package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(rearrangeArray(nums))
}

func rearrangeArray(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i, n := 0, len(nums)-1
	ans := make([]int, 0)
	for {
		if n <= i {
			break
		}
		ans = append(ans, nums[n])
		ans = append(ans, nums[i])
		i++
		n--
	}
	if n == i {
		ans = append(ans, nums[n])
	}
	return ans
}
