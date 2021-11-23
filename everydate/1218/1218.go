package main

func main() {
	serach()
}

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	count := 1
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > count {
			count = dp[v]
		}
	}
	return count
}

func serach() {
	chann := make(chan string)
	chann <- "weixin"
	chann <- "wx"
	//<-chann
}
