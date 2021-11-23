package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	head1 := l1
	//head2 := l1
	point := false
	sum := 0
	for l1 != nil && l2 != nil {
		if point {
			sum = l1.Val + l2.Val + 1
			point = false
		} else {
			sum = l1.Val + l2.Val
		}
		if sum > 9 {
			l1.Val = sum % 10
			point = true
		} else {
			l1.Val = sum
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		for l1 != nil {
			if point {
				sum = l1.Val + 1
				point = false
			} else {
				sum = l1.Val
			}
			if sum > 9 {
				l1.Val = sum % 10
				point = true
			} else {
				l1.Val = sum
			}
			l1 = l1.Next
		}
		if point {
			l1.Next = &ListNode{1, nil}
		}
		return reverseList(head1)
	}
	if l2 != nil {
		ll2 := l2
		for l2 != nil {
			if point {
				sum = l2.Val + 1
				point = false
			} else {
				sum = l2.Val
			}
			if sum > 9 {
				l2.Val = sum % 10
				point = true
			} else {
				l2.Val = sum
			}
			l2 = l2.Next
		}
		if point {
			l2.Next = &ListNode{1, nil}
		}
		head := reverseList(ll2)
		head.Next = reverseList(head1)
		return head1
	}
	if point {
		p := &ListNode{1, nil}
		p.Next = reverseList(head1)
		return p
	} else {
		return reverseList(head1)
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
