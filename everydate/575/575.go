package main

import "sort"

func main() {

}

func distributeCandies(candyType []int) int {
	sort.Slice(candyType, func(i, j int) bool {
		return candyType[i] <= candyType[j]
	})
	n := len(candyType)
	count := 1
	for i := 1; i < n; i++ {
		if candyType[i] != candyType[i-1] {
			count++
		}
	}
	if n/2 < count {
		return n / 2
	}
	return count
}
