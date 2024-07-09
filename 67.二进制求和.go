package cmd

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	reverse := func(str string) string {
		reversed := []rune(str)
		for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = reversed[j], reversed[i]
		}
		return string(reversed)
	}
	a = reverse(a)
	b = reverse(b)
	m, n := len(a), len(b)

	carry := 0
	i := 0

	var sb strings.Builder
	for i < max(m, n) || carry > 0 {
		val := carry
		if i < m {
			val += int(a[i] - '0')
		}
		if i < n {
			val += int(b[i] - '0')
		}
		sb.WriteString(strconv.Itoa(val % 2))
		carry = val / 2
		i++
	}

	return reverse(sb.String())
}

// @lc code=end
