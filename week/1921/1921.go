package main

import (
	"fmt"
	"sort"
)

func main() {
	dist := []int{4, 2, 3}
	speed := []int{2, 1, 1}
	fmt.Println(eliminateMaximum(dist, speed))
}

func eliminateMaximum(dist []int, speed []int) int {
	time := make([]int, 0)
	for i := 0; i < len(dist); i++ {
		time = append(time, (dist[i]-1)/(speed[i]))
	}
	sort.Ints(time)
	for i := 0; i < len(time); i++ {
		if time[i] < i {
			return i
		}
	}
	return len(time)
}
