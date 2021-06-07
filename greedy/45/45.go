package main

import (
	"fmt"
	"math"
)

func main() {
	slic := []int{2, 3,1,1,4}
	res := jump2(slic)
	fmt.Println(res)
}

//my function 有问题，有时间了好好改改
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var index int
	maxstep := math.MinInt8
	num := 0
	for i := 0; i < len(nums)-1; {
		//fmt.Println("for 1", num)
		//fmt.Println(i)
		
		maxstep = math.MinInt8
		for j := 0; j <= nums[i]; j++ {
			if j+i > len(nums) {
				return num + 1
			}
			if nums[j+i] > maxstep {
				maxstep = nums[j+i]
				index = j+i
			}
		}
		i = i + index
		num++
	}
	return num
}

//function 2
//一直保证end大于等于i
func jump2(nums []int) int{
	step := 0
	maxindex := 0
	end := 0
	i := 0
	for i=0; i<len(nums)-1; i++{
		maxindex = max(maxindex, i+nums[i])
		if i == end{
			end = maxindex
			step++
		}
	}
	return step
}

func max(i,j int) int {
	if i>j{
		return i
	}
	return j
}