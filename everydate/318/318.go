package main

import "strings"

func main() {

}

func maxProduct(words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[i], words[j]) && res < len(words[i])*len(words[j]) {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}

func maxProduct1(words []string) int {
	wordsbit := make([]int, len(words))
	word := ""
	res := 0
	for i := 0; i < len(words); i++ {
		word = words[i]
		for j := 0; j < len(word); j++ {
			wordsbit[i] |= 1 << (word[j] - 'a')
		}
	}

	for i := 0; i < len(wordsbit); i++ {
		for j := 0; j < len(wordsbit); j++ {
			if wordsbit[i]&wordsbit[j] == 0 && len(words[i])*len(words[j]) > res {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}
