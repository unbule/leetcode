package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "5F3Z-2e-9-w"
	k := 4
	fmt.Println(licenseKeyFormatting(s, k))
}

func licenseKeyFormatting(s string, k int) string {
	str := strings.ReplaceAll(s, "-", "")
	if len(str) == 0 {
		return ""
	}
	rem := len(str) % k
	res := []byte{}
	num := (len(str) - rem) / k
	if rem != 0 {
		res = append(res, str[:rem]...)
	}
	if len(res) != 0 {
		res = append(res, '-')
	}
	for i := 0; i < num; i++ {
		res = append(res, str[rem+i*k:rem+i*k+k]...)
		res = append(res, '-')
	}
	return strings.ToUpper(string(res[:len(res)-1]))
}
