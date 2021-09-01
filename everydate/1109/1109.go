package main

import "fmt"

func main() {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	fmt.Println(corpFlightBookings(bookings, 5))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n+1)
	for _, i := range bookings {
		for j := i[0]; j <= i[1]; j++ {
			ans[j] += i[2]
		}
	}
	return ans[1:]
}
