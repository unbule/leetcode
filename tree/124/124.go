package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	sumPath := math.MinInt32
	var gainsum func(*TreeNode) int

	gainsum = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		leftpath := max(gainsum(tn.Left), 0)
		rightpath := max(gainsum(tn.Right), 0)

		pricepath := tn.Val + leftpath + rightpath
		sumPath = max(pricepath, sumPath)

		return tn.Val + max(leftpath, rightpath)
	}
	gainsum(root)
	return sumPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
