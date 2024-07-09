package cmd

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	valToIndex := make(map[int]int)
	for k, v := range nums {
		need := target - v
		if valToIndex[need] != 0 {
			return []int{valToIndex[need] - 1, k}
		}
		valToIndex[v] = k + 1
	}
	return nil
}

// @lc code=end
