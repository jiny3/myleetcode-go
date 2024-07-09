package cmd

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	rightMap := make(map[int]int, 0)
	rightNodes(root, 0, &rightMap)
	res := []int{}
	for k, v := range rightMap {
		if len(res) <= k {
			res = append(res, make([]int, k-len(res)+1)...)
		}
		res[k] = v
	}
	return res
}

func rightNodes(root *TreeNode, height int, rightMap *map[int]int) {
	if root == nil {
		return
	}
	(*rightMap)[height] = root.Val
	rightNodes(root.Left, height+1, rightMap)
	rightNodes(root.Right, height+1, rightMap)
}

// @lc code=end
