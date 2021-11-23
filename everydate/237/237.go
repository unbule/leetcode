package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	p, q := node, node.Next
	for {
		p.Val = q.Val
		if q.Next == nil {
			break
		}
		p = q
		q = q.Next
	}
	p.Next = nil
}
