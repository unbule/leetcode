package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " Hello World "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	str := strings.Split(s, " ")
	return len(str[len(str)-1])
}
