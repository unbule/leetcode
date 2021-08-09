package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			slic := nums[j+1:]
			index := sort.SearchInts(slic, 0-(nums[i]+nums[j]))
			if index >= len(slic) {
				continue
			}
			if (nums[i]+nums[j])+slic[index] == 0 {
				cn := make([]int, 3)
				fmt.Println("numssss")
				cn[0] = nums[i]
				cn[1] = nums[j]
				cn[2] = slic[index]
				res = append(res, cn)
			}
		}
	}
	return res
}
