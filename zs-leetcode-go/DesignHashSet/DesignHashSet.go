package DesignHashSet

type MyHashSet struct {
	Data []*struct{}
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Data: make([]*struct{}, 1000000+1),
	}
}


func (this *MyHashSet) Add(key int)  {
	this.Data[key] = &struct{}{}
}


func (this *MyHashSet) Remove(key int)  {
	this.Data[key] = nil
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.Data[key] != nil
}
