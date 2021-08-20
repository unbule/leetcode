package main

import "fmt"

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrace(digits, 0, "")
	return combinations
}

func backtrace(digits string, index int, str string) {
	if index == len(digits) {
		combinations = append(combinations, str)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		letterscount := len(letters)
		for i := 0; i < letterscount; i++ {
			backtrace(digits, index+1, str+string(letters[i]))
		}
	}
}
