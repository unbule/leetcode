package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays2(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	ans := 0

	var backtrace func([]int, int)

	backtrace = func(path []int, index int) {
		fmt.Println(path)
		if len(path)%2 != 0 {
			for _, v := range path {
				ans = ans + v
			}
		}
		for i := index; i < n; i++ {
			path = append(path, arr[i])
			backtrace(path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrace([]int{}, 0)
	return ans
}

func sumOddLengthSubarrays1(arr []int) int {
	sum := 0
	for i := 1; i <= len(arr); i++ {
		if i%2 == 0 {
			continue
		}
		for l := 0; l+i <= len(arr); l++ {
			if l+i == len(arr) {
				sum += partSum(arr[l:])
			} else {
				sum += partSum(arr[l : l+i])
			}
		}
	}
	return sum
}
func partSum(slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return s
}

func sumOddLengthSubarrays2(arr []int) int {
	n := len(arr)
	ans := 0
	for i := 1; i < n; i++ {
		arr[i] = arr[i] + arr[i-1]
	}
	fmt.Println(arr)
	for i := 1; i <= n; i = i + 2 {
		left, right := 0, i
		ans = ans + arr[right-1]
		for right < n {
			ans = ans + arr[right] - arr[left]
			left++
			right++
		}
	}
	return ans
}
