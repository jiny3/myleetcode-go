package cmd

import "sort"

/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
type sumNode struct {
	Sum   int
	Count int
}

func averageOfLevels(root *TreeNode) []float64 {
	sumMap := make(map[int]sumNode, 0)
	sumHandler(root, 0, &sumMap)
	keys, res := []int{}, []float64{}
	for k := range sumMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		res = append(res, float64(sumMap[key].Sum)/float64(sumMap[key].Count))
	}
	return res
}

func sumHandler(root *TreeNode, height int, sumMap *map[int]sumNode) {
	if root == nil {
		return
	}
	if i, ok := (*sumMap)[height]; !ok {
		(*sumMap)[height] = sumNode{Sum: root.Val, Count: 1}
	} else {
		(*sumMap)[height] = sumNode{Sum: i.Sum + root.Val, Count: i.Count + 1}
	}
	sumHandler(root.Left, height+1, sumMap)
	sumHandler(root.Right, height+1, sumMap)
}

// @lc code=end
