package main

import (
	"strings"
)

func main() {

}

func findWords(words []string) []string {
	index := make(map[byte]int)
	res := make([]string, 0)
	index['q'] = 1
	index['w'] = 1
	index['e'] = 1
	index['r'] = 1
	index['t'] = 1
	index['y'] = 1
	index['u'] = 1
	index['i'] = 1
	index['o'] = 1
	index['p'] = 1
	index['a'] = 2
	index['s'] = 2
	index['d'] = 2
	index['f'] = 2
	index['g'] = 2
	index['h'] = 2
	index['j'] = 2
	index['k'] = 2
	index['l'] = 2
	index['z'] = 3
	index['x'] = 3
	index['c'] = 3
	index['v'] = 3
	index['b'] = 3
	index['n'] = 3
	index['m'] = 3

	for i := 0; i < len(words); i++ {
		boo := true
		word := strings.ToLower(words[i])
		flag := index[word[0]]
		for j := 0; j < len(word); j++ {
			if flag != index[word[j]] {
				boo = false
				break
			}
		}
		if boo == true {
			res = append(res, string(words[i]))
		}
	}
	return res
}
