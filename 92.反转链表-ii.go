package cmd

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	index := 1
	res := &ListNode{Next: head}
	pre := res
	for index < left {
		pre = pre.Next
		index++
	}
	cur := pre.Next
	tail := cur
	for index <= right {
		tmp := cur
		cur = cur.Next
		next := pre.Next
		pre.Next = tmp
		tmp.Next = next
		index++
	}
	tail.Next = cur

	return res.Next
}

// @lc code=end
