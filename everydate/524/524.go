package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abpcplea"
	dictionary := []string{"ale", "bpple", "apple", "monkey", "plea"}
	fmt.Println(findLongestWord(s, dictionary))
}

func findLongestWord(s string, dictionary []string) string {
	res := make([]string, 0)
	count := 0
	for i := 0; i < len(dictionary); i++ {
		if fit(s, dictionary[i]) {
			res = append(res, dictionary[i])
			count++
		}
	}
	if count > 1 {
		sort.Slice(res, func(i, j int) bool {
			if len(res[i]) == len(res[j]) {
				return res[i] < res[j]
			}
			return len(res[i]) > len(res[j])
		})
	}
	if len(res) == 0 {
		return ""
	}
	return res[0]
}

func fit(sorc string, dest string) bool {
	n := len(sorc)
	m := len(dest)
	i, j := 0, 0
	for i < n && j < m {
		if sorc[i] == dest[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j == m
}
