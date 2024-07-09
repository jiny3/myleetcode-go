package cmd

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if !equal(s[left], s[right]) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func equal(x byte, y byte) bool {
	if x >= 'A' && x <= 'Z' {
		x += 32
	}
	if y >= 'A' && y <= 'Z' {
		y += 32
	}
	fmt.Println(x, y)
	return x == y
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// @lc code=end
