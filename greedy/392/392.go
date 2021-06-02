/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
*/

package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "aabbcc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	for j := 0; j < len(t); j++ {
		if i < len(s) && s[i] == t[j] {
			i++
		}
	}
	if i == len(s) {
		return true
	}
	return false
}
