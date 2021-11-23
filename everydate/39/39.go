package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int, int, []int, []int)
	dfs = func(sum, index int, path []int, cand []int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := index; i < len(cand); i++ {
			if sum+cand[i] > target {
				continue
			}
			sum = sum + cand[i]
			path = append(path, cand[i])
			dfs(sum, i, path, cand)
			path = path[:len(path)-1]
			sum = sum - cand[i]
		}
	}
	dfs(0, 0, path, candidates)
	return res
}
