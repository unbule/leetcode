package main

import "fmt"

func main() {
	chalk := []int{5, 1, 5}
	k := 22
	fmt.Println(chalkReplacer(chalk, k))
}
func chalkReplacer(chalk []int, k int) int {
	total := 0
	for _, v := range chalk {
		total += v
	}
	k %= total
	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return 0
}
