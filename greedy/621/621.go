package main

import (
	"fmt"
	"sort"
)

func main() {
	tasks := "AAABBB"
	res := leastInterval([]byte(tasks), 2)
	fmt.Println(res)
}

// function 1 use sort.Sort to sort
type List []int

func (a List) Len() int           { return len(a) }
func (a List) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a List) Less(i, j int) bool { return a[i] > a[j] }

func leastInterval(tasks []byte, n int) int {
	max := 0
	taskMap := make(map[string]int)
	i := 0
	for i := 0; i < len(tasks); i++ {
		taskMap[string(tasks[i])] = taskMap[string(tasks[i])] + 1
	}
	taskList := make(List, len(taskMap))
	for _, v := range taskMap {
		taskList[i] = v
		i++
	}
	sort.Sort(taskList)
	for j := 0; j < len(taskList); j++ {
		if taskList[j] == taskList[0] {
			max = max + 1
		}
	}
	fmt.Println(taskList)
	fmt.Println(max)
	if (taskList[0]-1)*(n+1)+max < len(tasks) {
		return len(tasks)
	}
	return (taskList[0]-1)*(n+1) + max
}

//function 2
func leastInterval2(tasks []byte, n int) int {
    if n == 0 {
		return len(tasks)
	}
	list := make([]int, 26)
	for k := range tasks {
		list[tasks[k]-'A']++
	}
	maxExec, maxCount := 0, 0
	for k := range list {
		if list[k] > maxExec {
			maxExec, maxCount = list[k], 1
		} else if list[k] == maxExec {
			maxCount++
		}
	}
	time := (maxExec-1)*(n+1) + maxCount
	if time < len(tasks) {
		return len(tasks)
	}
	return time
}