package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	count := 1
	index := head
	for index.Next != nil {
		count++
		index = index.Next
	}
	fmt.Println(count)
	fk := count - k
	count = 1
	for count != fk {
		count++
		head = head.Next
	}
	return head
}
