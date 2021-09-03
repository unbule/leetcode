package main

import "fmt"

func main() {
	n := 8
	fmt.Println(wx(n))
}

func wx(n int) int {
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = i
	}
	count := 1
	j := 1
	ln := len(nums)
	//fmt.Println(nums)
	for len(nums) > 2 {
		count++
		j++
		if j >= ln {
			j = 1
		}
		if count == 3 {
			fmt.Println(j)
			copy(nums[j:], nums[j+1:])
			nums = nums[:len(nums)-1]
			ln = len(nums)
			if j == ln {
				j = 1
			}
			count = 1
			fmt.Println(nums)
		}
	}
	return nums[1]
}
