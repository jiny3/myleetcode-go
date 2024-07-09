package cmd

/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	h := 0
	hList := make([]int, n+1)
	for _, v := range citations {
		for i := 0; i <= min(n, v); i++ {
			hList[i]++
		}
	}
	for k, v := range hList {
		if v >= k {
			h = k
		}
	}
	return h
}

// @lc code=end
