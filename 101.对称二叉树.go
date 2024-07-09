package cmd

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	return isSame(root.Left, root.Right)
}

func isSame(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	if !isSame(left.Left, right.Right) {
		return false
	}
	return isSame(left.Right, right.Left)

}

// @lc code=end
