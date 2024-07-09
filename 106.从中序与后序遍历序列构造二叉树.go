package cmd

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
var inToIndex = make(map[int]int)

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	for k, v := range inorder {
		inToIndex[v] = k
	}
	return build(postorder, inorder, 0, n-1, 0, n-1)
}

func build(postorder, inorder []int, postStart, postEnd, inStart, inEnd int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{Val: postorder[postEnd]}
	index := inToIndex[postorder[postEnd]]

	root.Left = build(postorder, inorder, postStart, postStart+index-inStart-1, inStart, index-1)
	root.Right = build(postorder, inorder, postStart+index-inStart, postEnd-1, index+1, inEnd)
	return root
}

// @lc code=end
