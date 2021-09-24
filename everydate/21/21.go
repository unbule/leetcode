package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	q := head
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			q.Next = l1
			l1 = l1.Next
		} else {
			q.Next = l2
			l2 = l2.Next
		}
		q = q.Next
	}
	if l1 == nil {
		q.Next = l2
	} else {
		q.Next = l1
	}
	return head.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	lrehead := &ListNode{0, nil}
	lrev := lrehead
	for {
		if l1 == nil {
			break
		}
		if l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			lrev.Next = l1
			l1 = l1.Next
		} else {
			lrev.Next = l2
			l2 = l2.Next
		}
		lrev = lrev.Next
	}
	if l1 == nil {
		lrev.Next = l2
	} else {
		lrev.Next = l1
	}
	return lrehead.Next
}
