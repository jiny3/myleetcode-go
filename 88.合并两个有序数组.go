package cmd

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	mIndex, nIndex := m-1, n-1
	cur := len(nums1) - 1
	for mIndex >= 0 && nIndex >= 0 {
		if nums1[mIndex] > nums2[nIndex] {
			nums1[cur] = nums1[mIndex]
			mIndex--
		} else {
			nums1[cur] = nums2[nIndex]
			nIndex--
		}
		cur--
	}
	for nIndex >= 0 {
		nums1[cur] = nums2[nIndex]
		nIndex--
		cur--
	}
}

// @lc code=end
