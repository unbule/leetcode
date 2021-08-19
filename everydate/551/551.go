package main

import "fmt"

func main() {
	s := "PPALLL"
	fmt.Println(checkRecord(s))
}

func checkRecord(s string) bool {
	countA := 0
	countL := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countA++
			countL = 0
		}
		if s[i] == 'L' {
			countL++
		}
		if countL >= 3 {
			break
		}
		if s[i] == 'P' {
			countL = 0
		}
	}
	if countA < 2 && countL < 3 {
		return true
	}
	return false
}
