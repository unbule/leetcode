package main

import (
	"sort"
)

func main() {

}

func findLHS(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := 0
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
	}
	for k, v := range temp {
		if k1 := temp[k+1]; k1 > 0 && v+k1 > res {
			res = v + k1
		}
	}
	return res
}
