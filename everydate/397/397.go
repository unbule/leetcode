package main

func main() {

}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	if n%2 != 0 {
		return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
