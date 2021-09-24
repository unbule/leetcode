package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	seen := make([]bool, n)
	ans := make([][]int, 0)

	var backtrace func(int, []int)
	backtrace = func(index int, path []int) {
		if len(path) == n {
			res := make([]int, len(path))
			copy(res, path)
			ans = append(ans, res)
			return
		}
		for i := 0; i < n; i++ {
			fmt.Println(seen, i)
			//当前seen[i]为true
			//应填入nums[i]的下一个元素
			//seen[i]为false且seen[i-1]也为false
			//应填入nums[i]的上一个元素
			if seen[i] || i > 0 && !seen[i-1] && nums[i] == nums[i-1] {
				continue
			}
			seen[i] = true
			path = append(path, nums[i])
			backtrace(i+1, path)
			seen[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrace(0, []int{})
	return ans
}
