package cmd

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -9
	}
	return DP(coins, amount, dp)
}

func DP(coins []int, amount int, dp []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if dp[amount] != -9 {
		return dp[amount]
	}

	res := 100000
	for _, coin := range coins {
		temp := DP(coins, amount-coin, dp)
		if temp == -1 {
			continue
		}
		res = min(res, temp+1)
	}
	if res == 100000 {
		dp[amount] = -1
	} else {
		dp[amount] = res
	}
	return dp[amount]
}

// @lc code=end
