package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	p := head
	nhead := &ListNode{
		Next: nil,
	}
	for p != nil {
		node := &ListNode{
			Val: p.Val,
		}
		node.Next = nhead.Next
		nhead.Next = node
		p = p.Next
	}
	p = head
	q := nhead.Next
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}
