package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	version1 := "1.01"
	version2 := "1.001"
	fmt.Println(compareVersion(version1, version2))
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	n1, n2 := len(v1), len(v2)
	x, y := 0, 0
	for i := 0; i < n1 || i < n2; i++ {
		x, y = 0, 0
		if i < n1 {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < n2 {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
