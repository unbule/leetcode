package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(smallestK(arr, 4))
	fmt.Println(smallestK2(arr, 4))
}

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

type hp struct{ sort.IntSlice }

func (hp) Pop() (_ interface{}) { return }
func (h *hp) Push(interface{})  {}

//保证最大的元素出现在堆顶部
func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }

func smallestK2(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}
