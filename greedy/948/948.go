package greedy


func bagOfTokensScore(tokens []int, power int) int {
	right := len(tokens)-1
	left := 0
	sorce := 0
	for left <= right{
		for left<=right && power >= tokens[left] {
			power = power - tokens[left]
			sorce ++
			left ++
		}
		if left < right && power + tokens[right] >= tokens[left] && sorce > 0{
			power = power + tokens[right]
			sorce -- 
			right --
		}else {
			break
		}
	}
	return sorce
}