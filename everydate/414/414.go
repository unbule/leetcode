package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2, 2}
	fmt.Println(thirdMax(nums))
}

type hp struct {
	sort.IntSlice
}

func (hp) Pop() (_ interface{}) { return }
func (h hp) Push(x interface{}) {
	for i := 0; i < 3; i++ {
		if x.(int) == h.IntSlice[i] {
			return
		}
	}
	fmt.Println(h.IntSlice)
	if h.IntSlice[0] < x.(int) {
		h.IntSlice[0] = x.(int)
	}
}
func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func thirdMax(nums []int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	h := &hp{nums[:3]}
	heap.Init(h)
	for _, v := range nums[3:] {
		heap.Push(h, v)
	}
	if len(h.IntSlice) >= 3 {
		return h.IntSlice[2]
	}
	return h.IntSlice[0]
}
