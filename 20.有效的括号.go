package cmd

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}
	sLen := 0
	for _, v := range s {
		if strings.Contains("([{", string(v)) {
			stack = append(stack, byte(v))
			sLen++
		} else {
			if sLen == 0 || !isFix(stack[sLen-1], byte(v)) {
				return false
			}
			sLen--
			stack = stack[:sLen]
		}
	}
	return sLen == 0
}

func isFix(x, y byte) bool {
	return (x == '{' && y == '}') || (x == '[' && y == ']') || (x == '(' && y == ')')
}

// @lc code=end
