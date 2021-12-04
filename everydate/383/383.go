package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	slicran := make([]int, 26)
	slicmag := make([]int, 26)
	for i := 0; i < len(ransomNote); i++ {
		slicran[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		slicmag[magazine[i]-'a']++
	}
	for j := 0; j < len(slicran); j++ {
		if slicran[j] > slicmag[j] {
			return false
		}
	}
	return true
}
