package main

import "fmt"

func main() {
	n := 102
	fmt.Println(wx(n))
}

func wx(n int) int {
	ans := 0
	flag := true
	for i := 1; i <= n; i++ {
		flag = true
		point := make([]int, 10)
		num := i
		for num > 0 {
			point[num%10] += 1
			if point[num%10] > 1 {
				flag = false
				break
			}
			num = num / 10
		}
		if flag {
			ans++
		}
	}
	return ans
}
