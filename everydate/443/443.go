package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte{97, 98, 98, 99, 99, 99}
	fmt.Println(compress(chars))
}

func compress(chars []byte) int {
	if len(chars) < 2 {
		return 1
	}
	l, r := 0, 0
	index := 0
	num := 0
	for r < len(chars) {
		for r < len(chars) && chars[l] == chars[r] {
			r++
			num++
		}
		chars[index] = chars[l]
		index++
		charnum := strconv.Itoa(num)
		if num > 1 {
			for i := 0; i < len(charnum) && i < len(chars); i++ {
				chars[index] = charnum[i]
				index++
			}
		}
		l = r
		num = 0
	}
	return index
}
