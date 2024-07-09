package cmd

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Next: head}
	pre, left, right := res, head, head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right != nil {
		pre = pre.Next
		left = left.Next
		right = right.Next
	}
	pre.Next = left.Next

	return res.Next
}

// @lc code=end
