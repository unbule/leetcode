package main

func main() {

}

func arrangeCoins(n int) int {
	res := 0
	for i := 1; ; i++ {
		res += i
		if res > n {
			return i - 1
		}
	}
}
