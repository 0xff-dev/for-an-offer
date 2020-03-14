package LFUCache

import (
    "sort"
    "time"
)

type cacheType struct {
    Order time.Time
    AccessCnt int
    Key int
    Val int
}

type caches []*cacheType
func (self caches) Len() int {
    return len(self)
}
func (self caches) Less(i, j int) bool {
    if self[i].AccessCnt == self[j].AccessCnt {
        return self[i].Order.After(self[j].Order)
    }
    return self[i].AccessCnt > self[j].AccessCnt
}
func (self caches) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}


type LFUCache struct {
    Size int
    Index int
    Data caches
    DataMap map[int]*cacheType
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        Size: capacity,
        Index: 0,
        Data: make([]*cacheType, capacity),
        DataMap: map[int]*cacheType{},
    }
}


func (this *LFUCache) Get(key int) int {
    if this.Size == 0 {
        return -1
    }
    if val, ok := this.DataMap[key]; ok {
        this.DataMap[key].Order = time.Now()
        this.DataMap[key].AccessCnt++
        return val.Val
    } 
    return -1
}


func (this *LFUCache) Put(key int, value int)  {
    if this.Size == 0 { return }
    if _, ok := this.DataMap[key]; ok {
        this.DataMap[key].Order = time.Now()
        this.DataMap[key].Val = value
        this.DataMap[key].AccessCnt++
        return
    }
    newCacheData := &cacheType{
        Order: time.Now(),
        AccessCnt: 0,
        Key: key,
        Val: value,
    }
    this.DataMap[key] = newCacheData
    if this.Index < this.Size {
        this.Data[this.Index] = newCacheData
        this.Index++
    } else {
        sort.Sort(caches(this.Data))
        delete(this.DataMap, this.Data[this.Index-1].Key)
        this.Data[this.Index-1] = newCacheData
    }
}

