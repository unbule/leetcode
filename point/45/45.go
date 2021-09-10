package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber(nums))
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		str1 := strconv.Itoa(nums[i])
		str2 := strconv.Itoa(nums[j])
		return str1+str2 < str2+str1
	})
	ans := ""
	for i := 0; i < len(nums); i++ {
		ans += strconv.Itoa(nums[i])
	}
	return ans
}
