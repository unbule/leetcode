package main

import "strconv"

func main() {

}

func fizzBuzz(n int) []string {
	res := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			res = append(res, "Fizz")
			continue
		}
		if i%5 == 0 {
			res = append(res, "Buzz")
			continue
		}
		res = append(res, strconv.Itoa(i))
	}
	return res
}
