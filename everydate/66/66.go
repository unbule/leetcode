package main

func main() {

}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i]/10 == 0 {
			return digits
		}
		digits[i] = digits[i] % 10
	}
	return append([]int{1}, digits...)
}
