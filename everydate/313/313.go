package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	n := 5
	primes := []int{2, 7, 13, 19}
	nthSuperUglyNumber(n, primes)
}

type hp struct {
	sort.IntSlice
}

func nthSuperUglyNumber(n int, primes []int) int {
	seen := map[int]bool{1: true}
	ugly := 0
	h := &hp{[]int{1}}
	for i := 0; i < n; i++ {
		ugly = heap.Pop(h).(int)
		for _, p := range primes {
			if next := p * ugly; !seen[next] {
				seen[next] = true
				heap.Push(h, next)
			}
		}
	}
	return ugly
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
	fmt.Println("push: ", v.(int), h.IntSlice)
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	fmt.Println(a)
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	fmt.Println(v)
	return v
}
