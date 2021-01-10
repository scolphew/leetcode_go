package main

import (
	"fmt"
	"github.com/scolphew/leetcode_go/base"
)

type TreeNode = base.TreeNode

var prev, ans int

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	if tmp := node.Val - prev; tmp < ans {
		ans = tmp
	}
	prev = node.Val
	dfs(node.Right)
}

func minDiffInBST(root *TreeNode) int {
	ans = 0x7fffffff
	prev = -0x80000000
	dfs(root)
	return ans
}

func main() {
	x := minDiffInBST(&TreeNode{
		Val: 3,
	})

	fmt.Println(x)
}
