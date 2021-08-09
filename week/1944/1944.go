package main

import "fmt"

func main() {
	heights := []int{10, 6, 8, 5, 11, 9}
	fmt.Println(canSeePersonsCount(heights))
}

func canSeePersonsCount(heights []int) []int {
	res := make([]int, len(heights))
	count := 0
	for i := 0; i < len(heights)-1; i++ {
		for j := 1; j < len(heights); j++ {
			traget := heights[i]
			if traget > heights[j] && heights[j-1] < heights[j] {
				count++
			} else {
				break
			}
		}
		res[i] = count
		count = 0
	}
	return res
}
