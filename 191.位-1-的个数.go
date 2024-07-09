package cmd

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(n int) int {
	ans := 0
	for n != 0 {
		ans += (n & 1)
		n = n >> 1
	}
	return ans
}

// @lc code=end
