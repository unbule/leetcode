package main

func main() {

}

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for _, p := range points {
		mapp := map[int]int{}
		for _, q := range points {
			dis := (q[0]-p[0])*(q[0]-p[0]) + (q[1]-p[1])*(q[1]-p[1])
			mapp[dis]++
		}
		for _, m := range mapp {
			res += m * (m - 1)
		}
	}
	return res
}
