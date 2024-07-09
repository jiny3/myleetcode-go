package cmd

import "fmt"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	ans := []string{}
	pre := nums[0]
	cur := pre
	for i := 1; i < len(nums); i++ {
		if nums[i] == cur+1 {
			cur += 1
			continue
		}
		if cur == pre {
			ans = append(ans, fmt.Sprintf("%d", pre))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", pre, cur))
		}
		pre = nums[i]
		cur = pre
	}
	if cur == pre {
		ans = append(ans, fmt.Sprintf("%d", pre))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", pre, cur))
	}
	return ans
}

// @lc code=end
