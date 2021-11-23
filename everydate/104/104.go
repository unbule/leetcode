package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max, n := 0, 1
	var preorder func(*TreeNode, int)
	preorder = func(p *TreeNode, n int) {
		if p == nil {
			return
		}
		if max < n {
			max = n
		}
		preorder(p.Left, n+1)
		preorder(p.Right, n+1)
	}
	preorder(root, n)
	return max
}
