package cmd

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	count := len(nums) - (k % len(nums))

	for k, v := range append(nums[count:], nums[:count]...) {
		nums[k] = v
	}
}

// @lc code=end
