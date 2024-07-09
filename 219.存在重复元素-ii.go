package cmd

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := indexMap[v]; !ok {
			indexMap[v] = i
		} else {
			if (i - j) <= k {
				return true
			} else {
				indexMap[v] = i
			}
		}
	}
	return false
}

// @lc code=end
