package main

import "crypto/rand"

func main() {

}

type Solution struct {
	org, num []int
}

func Constructor(nums []int) Solution {
	return Solution{
		org: append([]int(nil), nums...),
		num: nums,
	}
}

func (this *Solution) Reset() []int {
	copy(this.num, this.org)
	return this.num
}

func (this *Solution) Shuffle() []int {
	res := make([]int, len(this.org))
	for i := range res {
		j := rand.Intn(len(this.num))
		res[i] = this.num[j]
		this.num = append(this.num[:j], this.num[j+1:]...)
	}
	this.num = res
	return this.num
}
