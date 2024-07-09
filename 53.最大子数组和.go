package cmd

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	cur := -10000
	res := nums[0]
	for _, v := range nums {
		cur = max(cur+v, v)
		res = max(res, cur)
	}
	return res
}

// @lc code=end
