package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	k := n / 3
	l, r := 0, 0
	count := 0
	index := 0
	for r <= n {
		if r == n {
			if count > k {
				nums[index] = nums[l]
				index++
				l = r
			}
			break
		}
		if nums[l] == nums[r] {
			count++
		} else if nums[l] != nums[r] {
			if count > k {
				nums[index] = nums[l]
				index++
				l = r
			} else {
				l = r
			}
			count = 1
		}
		r++
	}
	return nums[:index]
}
