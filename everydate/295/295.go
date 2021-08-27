package main

import (
	"fmt"
	"sort"
)

func main() {
	obj := Constructor()
	obj.AddNum(3)
	obj.AddNum(2)
	obj.AddNum(4)
	obj.AddNum(1)
	param_2 := obj.FindMedian()
	fmt.Println(param_2)
}

type MedianFinder struct {
	slic []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	n := len(this.slic)
	i := sort.Search(n, func(i int) bool {
		return this.slic[i] < num
	})
	if i == n {
		this.slic = append(this.slic, num)
	} else {
		this.slic = append(this.slic, 0)
		copy(this.slic[i+1:], this.slic[i:])
		this.slic[i] = num
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.slic)
	if n%2 == 0 {
		return (float64(this.slic[n/2]) + float64(this.slic[n/2-1])) / float64(2)
	}
	return float64(this.slic[n/2])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
