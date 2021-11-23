package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(2, 3))
}

// numerator/denominator

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	res := []byte{}
	if numerator < 0 != (denominator < 0) {
		res = append(res, '-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	zhengshu := numerator / denominator
	res = append(res, strconv.Itoa(zhengshu)...)
	res = append(res, '.')

	rem := map[int]int{}
	remainder := numerator % denominator
	for remainder != 0 && rem[remainder] == 0 {
		rem[remainder] = len(res)
		remainder *= 10
		res = append(res, '0'+byte(remainder/denominator))
		remainder = remainder % denominator
	}
	if remainder > 0 {
		index := rem[remainder]
		res = append(res[:index], append([]byte{'('}, res[index:]...)...)
		res = append(res, ')')
	}
	return string(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
