package main

import "fmt"

func main(){
	tasks := []string{"A","A","A","B","B","B"}
	leastInterval(tasks,2)
}

func leastInterval(tasks []string, n int) int {
	taskMap := make(map[string]int)
	i := 0
	for i:=0; i<len(tasks); i++{
		taskMap[string(tasks[i])] = taskMap[string(tasks[i])] + 1
	}
	for _, v := range taskMap{
		taskList[i] := v
		i++ 
	}
	fmt.Println(len(taskMap))
	return 0
}