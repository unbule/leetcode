package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := ""
	nums := make([]int, 0)
	fmt.Scan(&n)

	arr := strings.Split(n, ",")

	for i := 0; i < len(arr); i++ {
		num, _ := strconv.Atoi(arr[i])
		nums = append(nums, num)
	}
	k := 0
	prt := make([][]int, 0)
	for {
		res := make([]int, 0)
		for k = 0; k < len(nums); k++ {
			if nums[k] != 0 {
				res = append(res, nums[k])
				nums[k] = 0
				break
			}
		}
		if k == len(nums) {
			break
		}
		n := res[0]
		for i := k; i < len(nums); i++ {
			if nums[i] != 0 && n > nums[i] {
				res = append(res, nums[i])
				n = nums[i]
				nums[i] = 0
			}
		}
		prt = append(prt, res)
	}

	fmt.Println(len(prt))
	for i := 0; i < len(prt); i++ {
		for j := 0; j < len(prt[i]); j++ {
			if j+1 == len(prt[i]) {
				fmt.Printf("%d\n", prt[i][j])
			} else {
				fmt.Printf("%d,", prt[i][j])
			}
		}
	}
}
