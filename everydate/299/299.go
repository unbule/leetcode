package main

import "strconv"

func main() {

}

func getHint(secret string, guess string) string {
	secretslic := make([]int, 10)
	guessslic := make([]int, 10)
	// secretmap := make(map[int]byte)
	// for i := 0; i < len(secret); i++ {
	// 	secretmap[i] = secret[i]
	// }
	bull := 0
	cow := 0
	for j := 0; j < len(secret); j++ {
		secretslic[secret[j]-'0']++
	}

	for k := 0; k < len(guess); k++ {
		guessslic[guess[k]-'0']++
	}

	for x := 0; x < 10; x++ {
		cow += min(secretslic[x], guessslic[x])
	}

	for i := 0; i < len(guess); i++ {
		if secret[i] == guess[i] {
			bull++
		}
	}
	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow-bull) + "B"
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
