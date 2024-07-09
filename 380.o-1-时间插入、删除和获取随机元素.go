package cmd

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	mySet      []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mySet:      []int{},
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	this.valToIndex[val] = len(this.mySet)
	this.mySet = append(this.mySet, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	index := this.valToIndex[val]
	this.valToIndex[this.mySet[len(this.mySet)-1]] = index
	this.mySet[index], this.mySet[len(this.mySet)-1] = this.mySet[len(this.mySet)-1], this.mySet[index]
	this.mySet = this.mySet[:len(this.mySet)-1]
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.mySet[rand.Intn(len(this.mySet))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
