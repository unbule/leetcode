/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
*/
package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 1, 0, 4}
	res := canJump(a)
	fmt.Println(res)
}

// my function
// 从前向后
func canJump(nums []int) bool {
	sumStep := nums[0]
	i := 0
	if len(nums) == 1 {
		return true
	}
	if sumStep == 0 && len(nums) > 1 {
		return false
	}
	for i = 1; i < len(nums); i++ {
		sumStep = sumStep - 1
		if sumStep == 0 && nums[i] == 0 {
			break
		}
		if sumStep < nums[i] {
			sumStep = nums[i]
		}
	}
	if i >= len(nums)-1 {
		return true
	}
	return false
}

//function2 从后向前
func canJump2(nums []int) bool {
	step := 0
	j := len(nums) - 1
	for i:=len(nums)-2; i>=0; i--{
		step = nums[i] + i
		if step >= j{
			j = i
		}
	}
	return j<=0
}
