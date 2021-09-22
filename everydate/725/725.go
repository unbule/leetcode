package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	x, y := n/k, n%k
	res := make([]*ListNode, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		res[i] = curr
		size := x
		if i < y {
			size++
		}
		for j := 1; j < size; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return res
}
