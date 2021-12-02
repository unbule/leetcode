// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	l := Constructor(2)
// 	l.Put(1, 1)
// 	l.Put(2, 2)
// 	fmt.Println(l.Get(1))
// 	l.Put(3, 3)
// 	fmt.Println(l.Get(2))
// 	l.Put(4, 4)
// 	fmt.Println(l.Get(1))
// 	fmt.Println(l.Get(3))
// 	fmt.Println(l.Get(4))
// }

// type LRUCache struct {
// 	cap   int
// 	cache map[int]*ValTime
// }

// type ValTime struct {
// 	val  int
// 	time int
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		cap:   capacity,
// 		cache: make(map[int]*ValTime),
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	for i := range this.cache {
// 		if i == key {
// 			this.cache[i].time = 1
// 			continue
// 		}
// 		this.cache[i].time++
// 	}
// 	if _, ok := this.cache[key]; ok {
// 		if this.cache[key].val > 0 {
// 			return this.cache[key].val
// 		} else {
// 			return -1
// 		}
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {

// 	i := 0
// 	max := math.MinInt64
// 	del := 0
// 	for k := range this.cache {
// 		if this.cache[k].val > 0 {
// 			i++
// 		}
// 	}
// 	if i > this.cap {
// 		for key, val := range this.cache {
// 			if val.time > max {
// 				del = key
// 				max = val.time
// 			}
// 		}
// 		fmt.Println(1, this.cache[del].val)
// 		this.cache[del].val = -1
// 		this.cache[key] = &ValTime{
// 			time: 1,
// 			val:  value,
// 		}
// 		return
// 	} else {
// 		this.cache[key] = &ValTime{
// 			time: 1,
// 			val:  value,
// 		}
// 		return
// 	}
// }