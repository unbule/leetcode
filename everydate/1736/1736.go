package main

import "fmt"

func main() {
	s := "?7:?0"
	res := maximumTime(s)
	fmt.Println(res)
}

func maximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		if '4' <= t[1] && t[1] <= '9' {
			t[0] = '1'
		} else {
			t[0] = '2'
		}
	}
	if t[1] == '?' {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}
