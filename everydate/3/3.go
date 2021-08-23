package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abccdresssgffra"))
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	ans := 0
	l, r := 0, 0
	woidow := map[byte]int{}
	for r < n {
		if _, ok := woidow[s[r]]; !ok {
			woidow[s[r]] = r
		} else {
			if woidow[s[r]]+1 >= l {
				l = woidow[s[r]] + 1
			}
			woidow[s[r]] = r
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
