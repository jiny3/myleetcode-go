package cmd

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := n + 1
	curSum := 0
	left, right := 0, 0
	for left < n {
		if curSum < target {
			if right == n {
				break
			}
			curSum += nums[right]
			right++
		} else {
			res = min(res, right-left)
			curSum -= nums[left]
			left++
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}

// @lc code=end
