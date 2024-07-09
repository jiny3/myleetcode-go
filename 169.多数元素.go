package cmd

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	valCounter := make(map[int]int)
	maxCounter, maxVal := 0, nums[0]
	for _, v := range nums {
		valCounter[v]++
		if maxCounter < valCounter[v] {
			maxCounter = valCounter[v]
			maxVal = v
		}
	}
	return maxVal
}

// @lc code=end
