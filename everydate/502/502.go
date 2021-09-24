package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}

type pc struct {
	c, p int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(capital)
	arr := make([]pc, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, pc{p: profits[i], c: capital[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].c < arr[j].c
	})
	h := &hp{}
	for cur := 0; k > 0; k-- {
		for cur < n && arr[cur].c <= w {
			heap.Push(h, arr[cur].p)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w = w + heap.Pop(h).(int)
	}
	return w
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
