package main

type MyHashMap struct {
	dict []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	d := make([]int, 1000001)
	h := MyHashMap{d}
	return h
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if value != -1 {
		value += 1
	}
	this.dict[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	r := this.dict[key]
	if r != -1 {
		r -= 1
	}
	return r
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.dict[key] = 0
}
