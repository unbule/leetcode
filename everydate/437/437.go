package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, tarval int) (res int) {
	if root == nil {
		return
	}

	val := root.Val
	if root.Val == tarval {
		res++
	}

	res += rootSum(root.Left, tarval-val)
	res += rootSum(root.Right, tarval-val)

	return
}
