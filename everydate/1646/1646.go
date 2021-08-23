package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getMaximumGenerated(11))
}

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	ans := make([]int, n+1)
	ans[0] = 0
	ans[1] = 1
	for i := 1; 2*i <= n && 2*i+1 <= n; i++ {
		ans[i*2] = ans[i]
		ans[2*i+1] = ans[i] + ans[i+1]
		fmt.Println(ans[2*i])
	}
	fmt.Println(ans)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	return ans[0]
}
