package main

import (
	"fmt"
	"math/rand"
)

func main() {

}

type Solution struct {
	m, n, all int
	nums      map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m*n - 1, map[int]int{}}
}

func (this *Solution) Flip() []int {
	x := rand.Intn(this.all + 1)
	idx, ok := this.nums[x]
	if !ok {
		fmt.Println("111")
		idx = x
	}
	if v, okay := this.nums[this.all]; okay {
		this.nums[x] = v
	} else {
		this.nums[x] = this.all
	}
	this.all--
	return []int{idx / this.n, idx % this.n}
}

func (this *Solution) Reset() {
	this.all = this.m*this.n - 1
	this.nums = map[int]int{}
}
