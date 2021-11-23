package main

import "fmt"

func main() {
	s := "aaaaaaaaaaaaaaaaaaa"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	res := false
	var dfs func(string, []string)
	dfs = func(s string, wordDict []string) {
		if len(s) == 0 {
			res = true
			return
		}
		for i := 0; i < len(wordDict); i++ {
			wordlen := len(wordDict[i])
			if len(s)-wordlen >= 0 && s[len(s)-wordlen:] == wordDict[i] {
				s = s[:len(s)-wordlen]
				dfs(s, wordDict)
				s = s + wordDict[i]
			}
		}
	}
	dfs(s, wordDict)
	return res
}
