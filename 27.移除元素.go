package cmd

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	res := len(nums)
	index := len(nums) - 1
	for i := index; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[index] = nums[index], nums[i]
			res--
			index--
		}
	}
	return res
}

// @lc code=end
