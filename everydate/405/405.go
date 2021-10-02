package main

import (
	"fmt"
	"strings"
)

func main() {
	num := 0
	num = 26 >> 4
	fmt.Println(num)
	num = num >> 4
	fmt.Println(num)
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	st := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (i * 4) & 0xf
		if val > 0 || st.Len() > 0 {
			var d byte
			if val < 10 {
				d = '0' + byte(val)
			} else {
				d = 'a' + byte(val-10)
			}
			st.WriteByte(d)
		}
	}
	return st.String()
}
