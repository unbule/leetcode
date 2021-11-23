package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	sum := 0
	var after func(*TreeNode) int
	after = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		ls := after(tn.Left)
		rs := after(tn.Right)
		sum += abs(ls - rs)
		return ls + rs + tn.Val
	}
	after(root)
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
