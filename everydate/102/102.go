package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	temp := make(map[int][]int)
	res := make([][]int, 0)
	head := root
	pre(head, temp, 1)
	for _, v := range temp {
		res = append(res, v)
	}
	return res
}

func pre(p *TreeNode, temp map[int][]int, n int) {
	if p == nil {
		return
	}
	temp[n] = append(temp[n], p.Val)
	pre(p.Left, temp, n+1)
	pre(p.Right, temp, n+1)
}
