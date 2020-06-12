package insertdeletegetrandomo

import (
	"math/rand"
)

type RandomizedSet struct {
	locs map[int]int
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		locs: make(map[int]int),
		nums: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.locs[val]; ok {
		return false
	}
	this.locs[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	loc, ok := this.locs[val]
	if !ok {
		return false
	}
	lastone := this.nums[len(this.nums)-1]
	if val == lastone {
	} else {
		this.nums[loc] = lastone
		this.locs[lastone] = loc
	}
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.locs, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
