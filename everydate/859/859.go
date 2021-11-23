package main

func main() {

}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	index := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			index = append(index, i)
		}
	}
	if len(index) > 2 {
		return false
	}
	if len(index) == 2 {
		if (s[index[0]] == goal[index[1]] && s[index[1]] == goal[index[0]]) || (s[index[1]] == goal[index[0]] && s[index[0]] == goal[index[1]]) {
			return true
		} else {
			return false
		}
	}
	if len(index) == 0 {
		temp := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			temp[s[i]-'a']++
		}
		for _, v := range temp {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	return false
}
