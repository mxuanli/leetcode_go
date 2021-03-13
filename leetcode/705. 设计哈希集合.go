package main

type MyHashSet struct {
	dict []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	dict := make([]bool, 1000001)
	h := MyHashSet{dict: dict}
	return h
}

func (this *MyHashSet) Add(key int) {
	this.dict[key] = true
}

func (this *MyHashSet) Remove(key int) {
	if this.dict[key] == true {
		this.dict[key] = false
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.dict[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
