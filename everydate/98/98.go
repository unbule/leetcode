package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	max, min := math.MaxInt32, math.MinInt32
	var bl func(*TreeNode, int, int) bool
	bl = func(tn *TreeNode, max, min int) bool {
		if tn == nil {
			return true
		}
		if tn.Val <= min || tn.Val >= max {
			return false
		}
		return bl(tn.Left, tn.Val, min) && bl(tn.Right, max, tn.Val)
	}
	return bl(root, max, min)
}
