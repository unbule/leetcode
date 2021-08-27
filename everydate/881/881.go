package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []int{3, 2, 2, 1}
	limit := 3
	fmt.Println(numRescueBoats(people, limit))
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	light, weight := 0, len(people)-1
	ans := 0
	for light <= weight {
		if people[light]+people[weight] > limit {
			weight--
		} else {
			light++
			weight--
		}
		ans++
	}
	return ans
}
