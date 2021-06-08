package main

import (
	"fmt"
)

func main() {
	//bcabc
	//cbacdcbc
	s := "cbacdcbc"
	res := removeDuplicateLetters2(s)
	fmt.Println(res)
}

func removeDuplicateLetters(s string) string {
	liststres := make([]byte, 26)
	liststr := make([]int, 26)
	//str := ""
	//k := 1
	liststr[s[0]-'a'] = 1
	for i := 1 ; i < len(s); i++ {
		fmt.Println(liststr)
		if liststr[s[i]-'a'] == 0 {
			liststr[s[i]-'a'] = i + 1
			continue
		}
		if s[i] > s[i-1] && liststr[s[i]-'a'] < liststr[s[i-1]-'a']{
			liststr[s[i]-'a'] = i + 1
		}
	}
	//fmt.Println(liststr)
	for j := 0; j < len(liststr); j++ {
		if liststr[j] > 0 {
			liststres[liststr[j]] = byte(j + 'a')
		}
	}
	return string(liststres)
}


func removeDuplicateLetters2(s string) string{
	time := make([]int,26)
	var stack []byte
	inStack := make([]bool,26)
	for _, char := range s{
		time[char-'a']++
	}
	for i := range s{
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch <stack[len(stack)-1] && time[stack[len(stack)-1]-'a'] > 0{
				inStack[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		time[ch-'a']--
	}
	return string(stack)
}