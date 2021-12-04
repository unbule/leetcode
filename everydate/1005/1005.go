package main

import "sort"

func main() {

}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	m, ans := 101, 0
	for _, num := range nums {
		m = minAbs(m, num)
		if num < 0 && k > 0 {
			k--
			ans -= num
		} else {
			ans += num
		}
	}
	if k > 0 && k%2 == 1 {
		return ans - 2*m
	}
	return ans
}

func minAbs(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a > b {
		return b
	}
	return a
}
