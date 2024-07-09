package cmd

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix[0])
	for _, v := range matrix {
		if v[m-1] < target {
			continue
		} else if v[0] > target {
			return false
		}
		left, right := 0, m
		for left < right {
			mid := (left + right) / 2
			if v[mid] > target {
				right = mid
			} else if v[mid] < target {
				left = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// @lc code=end
