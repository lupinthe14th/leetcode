package designhashset

type MyHashSet struct {
	h []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{h: make([]bool, 1000000)}
}

func (this *MyHashSet) Add(key int) {
	this.h[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.h[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.h[key]
}
