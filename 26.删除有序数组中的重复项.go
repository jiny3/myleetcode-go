package cmd

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	res := len(nums)
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur] {
			cur++
			nums[cur] = nums[i]
		} else {
			res--
		}
	}
	return res
}

// @lc code=end
