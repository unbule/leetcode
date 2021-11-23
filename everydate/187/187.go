package main

func main() {

}

const L = 10

func findRepeatedDnaSequences(s string) []string {
	hmap := map[string]int{}
	res := []string{}
	for i := 0; i+L <= len(s); i++ {
		str := s[i : i+L]
		hmap[str]++
		if hmap[str] == 2 {
			res = append(res, str)
		}
	}
	return res
}
