package main

import "fmt"

func main() {
	s := "(){}[]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	n := len(s)
	str := make([]byte, 0)
	if n%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(str) == 0 || str[len(str)-1] != pairs[s[i]] {
				return false
			}
			str = str[:len(str)-1]
		} else {
			str = append(str, s[i])
		}
	}
	return len(str) == 0
}
