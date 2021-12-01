package main

import (
	"bytes"
)

func main() {

}

func originalDigits(s string) string {
	c := map[rune]int{}
	for _, v := range s {
		c[v]++
	}

	num := make([]int, 10)
	num[0] = c['z']
	num[2] = c['w']
	num[4] = c['u']
	num[6] = c['x']
	num[8] = c['g']

	num[1] = c['o'] - num[0] - num[2] - num[4]
	num[3] = c['h'] - num[8]
	num[5] = c['f'] - num[4]
	num[7] = c['s'] - num[6]
	num[9] = c['i'] - num[5] - num[6] - num[8]

	res := make([]byte, 0)
	for i := 0; i < 10; i++ {
		res = append(res, bytes.Repeat([]byte{byte('0' + i)}, num[i])...)
	}
	return string(res)
}
