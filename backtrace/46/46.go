package main

import "fmt"

func main() {
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	seen := make([]bool, len(nums))
	res := make([]int, 0)
	ans := make([][]int, 0)

	var backtrace func([]bool, []int)
	backtrace = func(seen []bool, res []int) {
		//fmt.Println(res)
		if len(res) == len(nums) {
			slic := make([]int, len(res))
			copy(slic, res)
			ans = append(ans, slic)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !seen[i] {
				seen[i] = true
				res = append(res, nums[i])
				backtrace(seen, res)
				res = res[:len(res)-1]
				seen[i] = false
			}
		}
	}
	backtrace(seen, res)
	return ans
}
