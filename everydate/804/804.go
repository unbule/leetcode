package main

import "fmt"

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
	mos := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	mapp := make(map[string]int)
	//mid := make(map[string]int)
	str := ""
	for i := 0; i < len(words); i++ {
		for _, ch := range words[i] {
			str = str + mos[ch-97]
		}
		mapp[str]++
		str = ""
	}
	return len(mapp)
}
