package cmd

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	n := len(s)
	StoT := make(map[byte]byte)
	TtoS := make(map[byte]byte)
	for i := 0; i < n; i++ {
		if _, ok := StoT[s[i]]; !ok {
			StoT[s[i]] = t[i]
		} else if StoT[s[i]] != t[i] {
			return false
		}
		if _, ok := TtoS[t[i]]; !ok {
			TtoS[t[i]] = s[i]
		} else if TtoS[t[i]] != s[i] {
			return false
		}
	}
	return true
}

// @lc code=end
