package main

func main() {

}

type SummaryRanges struct {
	nums      [10001]int
	beginStop map[int]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		beginStop: make(map[int]int),
	}
}

func (this *SummaryRanges) AddNum(val int) {
	pre := val - 1
	for ; pre >= 0; pre-- {
		if this.nums[pre] == 0 {
			break
		}
	}

	last := val + 1
	for ; last < len(this.nums); last++ {
		if this.nums[last] == 0 {
			break
		}
	}

	this.nums[val]++

	if last != val+1 {
		delete(this.beginStop, val+1)
	}
	this.beginStop[pre+1] = last - 1
}

func (this *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, 0, len(this.beginStop))

	for i := 0; i < len(this.nums); i++ {
		if val, ok := this.beginStop[i]; ok {
			res = append(res, []int{i, val})
			i = val
		}
	}
	return res
}
