package main

import "fmt"

func main() {
	numWaterBottles(15, 4)
}

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		fmt.Println(numBottles)
		res += numBottles / numExchange
		temp := numBottles - (numBottles/numExchange)*numExchange
		numBottles = numBottles/numExchange + temp
	}
	return res
}
