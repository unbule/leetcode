package main

import "fmt"

func main() {
	print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := getLength(head)
	d := len - n
	headd := &ListNode{0, head}
	p := headd
	for i := 0; i < d; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return headd.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexWaiterShift = iota
)

func print() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexWaiterShift)
}
