package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(subrob(nums[:n-1]), subrob(nums[1:]))
}

func max(i, j int) int {
	if i > j {
		fmt.Println("max", i)
		return i
	}
	return j
}

func subrob(nums []int) int {
	f, s := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		fmt.Println(v)
		f, s = s, max(f+v, s)
	}
	return s
}
