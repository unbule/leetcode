package main

import "fmt"

func partition(array []int, i int, j int) int {
	//第一次调用使用数组的第一个元素当作基准元素
	pivot := array[i]
	for i < j {
		for j > i && array[j] > pivot {
			j--
		}
		if j > i {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < pivot {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = pivot
	return i
}

func quicksort(array []int, low int, high int) {
	var pivotPos int //划分基准元素索引
	if low < high {
		pivotPos = partition(array, low, high)
		quicksort(array, low, pivotPos-1)
		quicksort(array, pivotPos+1, high)
	}
}

func main() {
	var arr = []int{56, 34, 58, 26, 79, 64, 37}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
