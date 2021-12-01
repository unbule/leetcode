package main

import "fmt"

func main() {
	slic := []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}
	fmt.Println(dailyTemperatures(slic))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	//i, j := 0, 0
	count := 1
	for i := 0; i < len(temperatures); i++ {
		count = 1
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				res[i] = count
				//count = 1
				break
			} else {
				count++
			}
		}
	}
	return res
}
