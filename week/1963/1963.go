package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "]]][[["
	fmt.Println(minSwaps(str))
}

func minSwaps(s string) int {
	count := 0
	res := 0
	cnt := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			count++
			cnt = append(cnt, count)
		}
		if s[i] == ']' {
			count--
			cnt = append(cnt, count)
		}
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] <= cnt[j]
	})
	if -cnt[0]%2 == 0 {
		res = -cnt[0] / 2
	}
	if -cnt[0]%2 != 0 {
		res = -cnt[0]/2 + 1
	}
	return res
}
