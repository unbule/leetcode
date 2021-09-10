package main

import "fmt"

func main() {
	pushed := []int{1, 0}
	popped := []int{1, 0}
	fmt.Println(validateStackSequences(pushed, popped))
}

type stack struct {
	nums []int
}

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	j := 0
	st := stack{nums: []int{}}
	for i := 0; i < n; i++ {
		st.push(pushed[i])
		//fmt.Println("push", st.nums)
		for len(st.nums) != 0 && j < n && st.nums[len(st.nums)-1] == popped[j] {
			st.pop()
			//fmt.Println("pop", st.nums)
			j++
		}
	}
	if len(st.nums) == 0 && j == n {
		return true
	}
	return false
}

func (st *stack) push(num int) {
	st.nums = append(st.nums, num)
}

func (st *stack) pop() int {
	a := st.nums
	res := a[len(a)-1]
	st.nums = a[:len(a)-1]
	return res
}
