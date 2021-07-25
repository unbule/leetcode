package main

import (
	"fmt"
	"sort"
)

func main() {
	s := [][]int{{1, 4}, {2, 3}, {4, 6}}
	fmt.Println(smallestChair(s, 1))
}

func smallestChair(times [][]int, targetFriend int) int {
	resfriend := times[targetFriend][0]

	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	fmt.Println(times)
	traget := 0
	for traget = 0; traget < len(times); traget++ {
		if resfriend == times[traget][0] {
			break
		}
	}
	seats := make([]int, len(times))
	res := -1
	for i := 0; i <= traget; i++ {
		for j := 0; j < len(seats); j++ {
			if times[i][0] >= seats[j] {
				seats[j] = times[i][1]
				res = j
				break
			}
			if times[i][0] < seats[j] {
				continue
			}
		}
	}
	return res
}
