package cmd

import "math"

/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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

func getMinimumDifference(root *TreeNode) int {
	prev := (*TreeNode)(nil)
	res := math.MaxInt32

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}

		prev = root
		traverse(root.Right)
	}

	traverse(root)
	return res
}

// @lc code=end
