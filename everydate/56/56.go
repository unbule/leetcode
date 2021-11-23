package main

import "sort"

func main() {

}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	ix, iy := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if iy >= intervals[i][0] {
			if iy <= intervals[i][1] {
				iy = intervals[i][1]
			}
		} else {
			res = append(res, []int{ix, iy})
			ix = intervals[i][0]
			iy = intervals[i][1]
		}
	}
	res = append(res, []int{ix, iy})
	return res
}
