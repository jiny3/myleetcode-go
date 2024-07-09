package cmd

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0
	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
			tIndex++
			continue
		}
		tIndex++
	}
	return sIndex == len(s)
}

// @lc code=end
