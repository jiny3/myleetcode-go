package cmd

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, v := range magazine {
		magazineMap[v]++
	}
	for _, v := range ransomNote {
		if magazineMap[v] == 0 {
			return false
		}
		magazineMap[v]--
	}
	return true
}

// @lc code=end
