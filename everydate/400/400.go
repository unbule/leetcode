package main

import "fmt"

func main() {
	fmt.Println(findNthDigit(11))
}

func findNthDigit(n int) int {
	i := 1
	count := 0
	res := 0
	for {
		num := i
		fmt.Println(num)
		for num > 0 {
			res = num % 10
			num = num / 10
			count++
			if count == n {
				return res
			}
		}
		i++
	}
}
