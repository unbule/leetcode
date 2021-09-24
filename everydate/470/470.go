package main

import "fmt"

func main() {
	fmt.Println(rand10())
}

/**
**随机数生成器
**(randX - 1) * Y + randY ==> 可以等概率的生成[1, X * Y]范围的随机数
**randX % X + 1 ==> 可以等概率的生成[1, X]范围的随机数
**/

func rand10() int {
	for {
		in := (rand7()-1)*7 + rand7()
		if in <= 40 {
			return in%10 + 1
		}
	}
}

func rand7() int {
	return 4
}
