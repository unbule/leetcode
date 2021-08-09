package main

import (
	"fmt"
)

func main() {
	rains := []int{1, 2, 3, 4}
	fmt.Println(avoidFlood(rains))
}

func avoidFlood(rains []int) []int {
	ans := make([]int, len(rains))
	nums := make([][]int, 0)
	m := make(map[int]int, 0)
	isFlood := false
	j := -1
	num := make([]int, 0)
	for i := 0; i < len(rains); i++ {
		if rains[i] > 0 {
			isFlood = true
		} else {
			if isFlood {
				if j > -1 {
					nums = append(nums, num)
				}
				num = make([]int, 0)
				j++
				isFlood = false
			}
			if j > -1 {
				num = append(num, i)
			}
		}
	}
	if len(num) > 0 {
		nums = append(nums, num)
	}
	j = 0
	isFlood = false
	for i := 0; i < len(rains); i++ {
		if rains[i] > 0 {
			isFlood = true
			ans[i] = -1
			if v, ok := m[rains[i]]; ok {
				if v == j {
					return []int{}
				}
				t := true
				for k := v; k < len(nums); k++ {
					if len(nums[k]) > 0 {
						ans[nums[k][0]] = rains[i]
						nums[k] = nums[k][1:]
						t = false
						break
					}
				}
				if t {
					return []int{}
				}
			}
			m[rains[i]] = j
		} else {
			if isFlood {
				j++
				isFlood = false
			}
			ans[i] = 1
		}
	}
	return ans
}
