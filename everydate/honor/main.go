package main

import (
	"fmt"
	"sort"
)

type Job struct {
	furit  string
	id     int
	weight int
}

func main() {
	var furit string
	var people, id, weight int
	jobs := make([]Job, 0)
	fmt.Scanf("%d\n", &people)

	for i := 0; i < 8; i++ {
		var job Job
		fmt.Scanln("%s %d %d\n", &furit, &id, &weight)
		job.furit = furit
		job.id = id
		job.weight = weight
		jobs = append(jobs, job)
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].furit < jobs[j].furit
	})
	j := 0
	for i := 1; i < len(jobs); i++ {
		if jobs[i].furit != jobs[i-1].furit {
			sort.Slice(jobs[j:i], func(i, j int) bool {
				return jobs[i].weight > jobs[j].weight
			})
			j = i
		}
	}
}
