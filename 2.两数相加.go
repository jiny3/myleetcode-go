package cmd

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := 0
	resNode := &ListNode{}
	preNode := resNode
	for l1 != nil || l2 != nil || res > 0 {
		if l1 != nil {
			res += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			res += l2.Val
			l2 = l2.Next
		}
		preNode.Next = &ListNode{Val: res % 10}
		res /= 10
		preNode = preNode.Next
	}
	return resNode.Next
}

// @lc code=end
