package main

import (
	"container/heap"
	"sort"
)

func scheduleCourse2(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][i]
	})
	h := &Heap{}
	total := 0
	for _, cours := range courses {
		if t := cours[0]; cours[1] >= t+total {
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = h.IntSlice[:len(a)-1]
	return x
}
