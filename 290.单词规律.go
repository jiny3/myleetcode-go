package cmd

import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	n := len(pattern)
	sArray := strings.Split(s, " ")
	patternToSArray := make(map[byte]string)
	sArrayToPattern := make(map[string]byte)

	if n != len(sArray) {
		return false
	}
	for i := 0; i < n; i++ {
		if _, ok := patternToSArray[pattern[i]]; !ok {
			patternToSArray[pattern[i]] = sArray[i]
		} else if patternToSArray[pattern[i]] != sArray[i] {
			return false
		}
		if _, ok := sArrayToPattern[sArray[i]]; !ok {
			sArrayToPattern[sArray[i]] = pattern[i]
		} else if sArrayToPattern[sArray[i]] != pattern[i] {
			return false
		}
	}
	return true
}

// @lc code=end
