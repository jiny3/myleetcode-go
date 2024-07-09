package cmd

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

var valToIndex = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	for i, v := range inorder {
		valToIndex[v] = i
	}
	return build(preorder, inorder, 0, n-1, 0, n-1)
}

func build(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[preStart]}
	index := valToIndex[preorder[preStart]]

	root.Left = build(preorder, inorder, preStart+1, preStart+index-inStart, inStart, index-1)
	root.Right = build(preorder, inorder, preStart+index-inStart+1, preEnd, index+1, inEnd)
	return root
}

// @lc code=end
