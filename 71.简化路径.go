package cmd

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := ""
	res := []string{}
	for _, v := range path {
		if byte(v) != '/' {
			stack += string(v)
			continue
		}
		switch stack {
		case "":
			continue
		case ".":
			stack = ""
		case "..":
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			stack = ""
		default:
			res = append(res, stack)
			stack = ""
		}
	}
	switch stack {
	case "":
	case ".":
	case "..":
		if len(res) > 0 {
			res = res[:len(res)-1]
		}
	default:
		res = append(res, stack)
	}
	return "/" + strings.Join(res, "/")
}

// @lc code=end
