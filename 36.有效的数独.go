package cmd

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]bool
	var boxs [3][3][9]bool
	for i, rv := range board {
		for j, cv := range rv {
			if cv == '.' {
				continue
			}
			index := cv - '1'
			if rows[i][index] || columns[j][index] || boxs[i/3][j/3][index] {
				return false
			}
			rows[i][index] = true
			columns[j][index] = true
			boxs[i/3][j/3][index] = true
		}
	}
	return true
}

// @lc code=end
