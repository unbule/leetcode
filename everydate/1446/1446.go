package main

import "fmt"

func main() {
	fmt.Println(maxPower("syui"))
}

func maxPower(s string) int {
	if len(s) == 1 {
		return 1
	}
	res := 0
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if res < count {
				res = count
			}
		} else {
			if res < count {
				res = count
			}
			count = 1
		}
	}
	return res
}
