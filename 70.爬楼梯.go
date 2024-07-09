package cmd

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	pre := 1
	cur := 1
	for i := 2; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

// @lc code=end
