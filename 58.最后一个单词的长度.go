package cmd

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	ans := 0
	cur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if cur != 0 {
				ans = cur
			}
			cur = 0
		} else {
			cur++
		}
	}
	if cur != 0 {
		ans = cur
	}
	return ans
}

// @lc code=end
