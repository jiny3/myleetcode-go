package cmd

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	var pre rune
	res := 0
	for _, v := range s {
		switch v {
		case 'I':
			res += 1
		case 'V':
			res += 5
			if pre == 'I' {
				res -= 2
			}
		case 'X':
			res += 10
			if pre == 'I' {
				res -= 2
			}
		case 'L':
			res += 50
			if pre == 'X' {
				res -= 20
			}
		case 'C':
			res += 100
			if pre == 'X' {
				res -= 20
			}
		case 'D':
			res += 500
			if pre == 'C' {
				res -= 200
			}
		case 'M':
			res += 1000
			if pre == 'C' {
				res -= 200
			}
		}
		pre = v
	}
	return res
}

// @lc code=end
