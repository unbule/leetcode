package main

func main() {

}

func subArrayRanges(nums []int) int64 {
	var ans int64 = 0
	for i, num := range nums {
		max, min := num, num
		for _, v := range nums[i+1:] {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
			ans += int64(max - min)
		}
	}
	return ans
}
