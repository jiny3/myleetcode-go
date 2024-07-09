package cmd

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	end, farthest, jumps := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}

// @lc code=end
