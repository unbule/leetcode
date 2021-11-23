package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(*Node, int)
	dfs = func(n *Node, i int) {
		if n.Children == nil {
			return
		}
		for _, node := range n.Children {
			i = i + 1
			if res < i {
				res = i
			}
			dfs(node, i)
			i = i - 1
		}
	}
	dfs(root, 0)
	return res
}
