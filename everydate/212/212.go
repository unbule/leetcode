package main

import "fmt"

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}

type Trie struct {
	children [26]*Trie
	word     string
}

func (tr *Trie) insert(word string) {
	node := tr
	for _, d := range word {
		d -= 'a'
		if node.children[d] == nil {
			node.children[d] = &Trie{}
		}
		node = node.children[d]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	tr := &Trie{}
	for _, word := range words {
		tr.insert(word)
	}
	m, n := len(board), len(board[0])
	seen := map[string]bool{}

	var traceback func(node *Trie, x, y int)
	traceback = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.children[ch-'a']
		if node == nil {
			return
		}

		if node.word != "" {
			seen[node.word] = true
		}

		board[x][y] = '#'

		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
				traceback(node, nx, ny)
			}
		}
		board[x][y] = ch
	}
	for i, row := range board {
		for j := range row {
			traceback(tr, i, j)
		}
	}
	ans := make([]string, 0, len(seen))
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}
