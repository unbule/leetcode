package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := []int{5, 4, 3, 2, 1}
	findRelativeRanks(s)
}

func findRelativeRanks(score []int) []string {
	temp := make(map[int]int)
	src := make([]int, len(score))
	copy(src, score)
	fmt.Println(src)
	res := make([]string, 0)
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})
	for i := 0; i < len(score); i++ {
		temp[score[i]] = i + 1
	}
	for i := 0; i < len(src); i++ {
		if temp[src[i]] == 1 {
			res = append(res, "Gold Medal")
		} else if temp[src[i]] == 2 {
			res = append(res, "Silver Medal")
		} else if temp[src[i]] == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(temp[src[i]]))
		}
	}
	return res
}
