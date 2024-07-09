package cmd

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	res := 0
	Len := len(prices)
	minVals, maxVals := make([]int, Len), make([]int, Len)
	minVals[0] = prices[0]
	maxVals[Len-1] = prices[Len-1]
	for i := 1; i < Len; i++ {
		minVals[i] = min(minVals[i-1], prices[i])
		maxVals[Len-i-1] = max(maxVals[Len-i], prices[Len-i-1])
	}
	// minVals[0] = prices[0]
	// maxVals[Len-1] = prices[Len-1]
	for i := 0; i < Len; i++ {
		res = max(res, maxVals[i]-minVals[i])
	}
	return res
}

// @lc code=end
