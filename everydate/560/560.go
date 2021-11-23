package main

import "fmt"

func main() {
	nums := []int{1, -1, 0}
	fmt.Println(subarraySum(nums, 0))
}

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	i, j := 0, 0
	for i = 0; i < len(nums); i++ {
		sum = 0
		for j = i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
