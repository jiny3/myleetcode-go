package cmd

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 随机链表的复制
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	toCopyNode := make(map[*Node]*Node, 0)
	res := &Node{}
	resCur := res
	for cur := head; cur != nil; cur = cur.Next {
		resCur.Next = &Node{Val: cur.Val}
		resCur = resCur.Next
		toCopyNode[cur] = resCur
	}
	resCur = res
	for cur := head; cur != nil; cur = cur.Next {
		resCur.Next.Random = toCopyNode[cur.Random]
		resCur = resCur.Next
	}
	return res.Next
}

// @lc code=end
