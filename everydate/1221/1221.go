package main

import "fmt"

func main() {
	s := "LLLLRRRR"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	countR := 0
	countL := 0
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			countL++
		}
		if s[i] == 'R' {
			countR++
		}
		if countL == countR {
			ans++
			countL = 0
			countR = 0
		}
	}
	return ans
}

func balancedStringSplit2(s string) int {
	count := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			count++
		}
		if s[i] == 'R' {
			count--
		}
		if count == 0 {
			ans++
		}
	}
	return ans
}
