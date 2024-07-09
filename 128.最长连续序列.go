package cmd

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	for _, v := range nums {
		numMap[v] = true
	}
	res := 0
	for k := range numMap {
		if numMap[k-1] {
			continue
		}
		curNum := k
		curLen := 1

		for numMap[curNum+1] {
			curNum += 1
			curLen += 1
		}
		res = max(res, curLen)
	}
	return res
}

// @lc code=end
