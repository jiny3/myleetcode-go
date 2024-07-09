package cmd

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x <= 0 {
		return x == 0
	}
	temp := x
	y := 0

	for temp > 0 {
		lastNum := temp % 10
		temp /= 10
		y = y*10 + lastNum
	}

	// 如果 x 和 y 相等，那么 x 就是回文数
	return y == x
}

// @lc code=end
