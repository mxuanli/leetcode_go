package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool           {}
func (this NestedInteger) GetInteger() int           {}
func (n *NestedInteger) SetInteger(value int)        {}
func (this NestedInteger) GetList() []*NestedInteger {}

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var stack []*NestedInteger
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	n := &NestedIterator{stack: stack}
	return n
}

func (this *NestedIterator) Next() int {
	r := this.stack[len(this.stack)-1].GetInteger()
	this.stack = this.stack[:len(this.stack)-2]
	return r
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		cur := this.stack[len(this.stack)-1]
		if cur.IsInteger() {
			return true
		} else {
			this.stack = this.stack[:len(this.stack)-1]
			list := cur.GetList()
			for i := len(list) - 1; i >= 0; i-- {
				this.stack = append(this.stack, list[i])
			}
		}
	}
	return false
}
