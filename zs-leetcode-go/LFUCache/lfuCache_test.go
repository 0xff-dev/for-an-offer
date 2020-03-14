package LFUCache

import (
    "fmt"
    "testing"
)

func TestLFUCache(t *testing.T) {
   obj := Constructor(2)
   obj.Put(1, 1)
   obj.Put(2,2)
   fmt.Println("get: ", obj.Get(1))
   obj.Put(3, 3)
   fmt.Println("get: ", obj.Get(2))
   fmt.Println("get: ", obj.Get(3))
   obj.Put(4, 4)
   fmt.Println("get: ", obj.Get(1))
   fmt.Println("get: ", obj.Get(3))
   fmt.Println("get: ", obj.Get(4))
}
