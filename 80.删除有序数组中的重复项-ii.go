package cmd

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	res := len(nums)
	pre, cur := 0, 1
	for i := 2; i < len(nums); i++ {
		if nums[cur] != nums[i] {
			pre, cur = pre+1, cur+1
			nums[cur] = nums[i]
		} else {
			if nums[cur] != nums[pre] {
				pre, cur = pre+1, cur+1
				nums[cur] = nums[i]
			} else {
				res--
			}
		}
	}
	return res
}

// @lc code=end
