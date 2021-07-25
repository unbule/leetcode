package main

import (
	"fmt"
)

func main() {
	s := "oa" //  12552031545
	fmt.Println(getLucky(s, 1))
	//fmt.Println(getLucky1(s, 5))
}

func getLucky(s string, k int) int {
	nslic := make([]int, len(s))
	for j := 0; j < len(s); j++ {
		nslic[j] = int(s[j] - 96)
	}
	num := nslic[0]
	x := 0
	for i := 1; i < len(nslic); i++ {
		x = 10
		for x <= nslic[i] {
			x = x * 10
		}
		num = num*x + nslic[i]
	}
	fmt.Println(num)
	res := 0
	for s := 0; s < k; s++ {
		for num > 0 {
			res = res + num%10 //198 19 1
			num = num / 10
		}
		num = res
		//fmt.Println(num)
		res = 0
	}
	return num
}

func getLucky1(s string, k int) (sum int) {
	for _, b := range s {
		b -= 'a' - 1
		sum += int(b/10 + b%10)
	}
	fmt.Println(sum)
	for k--; k > 0 && sum > 9; k-- {
		s := sum
		for sum = 0; s > 0; s /= 10 {
			sum += s % 10
		}
	}
	return
}
