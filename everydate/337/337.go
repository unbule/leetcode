package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	Select := node.Val + l[1] + r[1]
	NoSelect := max(l[0], l[1]) + max(r[0], r[1])
	return []int{Select, NoSelect}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
