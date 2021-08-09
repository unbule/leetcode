package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	piles := []int{4, 3, 6, 7}
	fmt.Println(minStoneSum(piles, 3))
}

type hp struct {
	sort.IntSlice
}

func minStoneSum(piles []int, k int) int {
	h := &hp{piles}
	sum := 0
	heap.Init(h)
	for ; k > 0; k-- {
		fmt.Println(h.IntSlice)
		fmt.Println("pop", h.IntSlice[0]/2)
		h.IntSlice[0] = h.IntSlice[0] - h.IntSlice[0]/2
		heap.Fix(h, 0)
	}
	for _, v := range h.IntSlice {
		sum += v
	}
	return sum
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (hp) Push(interface{}) {
}

func (h hp) Pop() (_ interface{}) {
	return
}
