package main

import (
	"strings"
)

func main() {

}

func truncateSentence(s string, k int) string {
	temp := strings.Split(s, " ")
	res := ""
	for i := 0; i < k; i++ {
		res += temp[i]
		if i == k-1 {
			continue
		}
		res += " "
	}
	return res
}
