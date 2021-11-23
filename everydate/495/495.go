package main

import "fmt"

type t struct {
	name int
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] >= duration {
			res += duration
		} else {
			res += timeSeries[i] - timeSeries[i-1]
		}
	}
	return res + duration
}

func main() {
	f := func() { fmt.Print("A") }
	defer f()
	f = func() { fmt.Print("B") }
	defer f()
}
