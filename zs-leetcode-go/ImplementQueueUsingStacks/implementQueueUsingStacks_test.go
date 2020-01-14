package ImplementQueueUsingStacks

import (
    "fmt"
    "testing"
)

func TestImplementQueueUsingStacks(t *testing.T) {
    obj := Constructor()
    obj.Push(1)
    obj.Push(2)
    fmt.Println(obj.Peek())
    fmt.Println(obj.Peek())
    fmt.Println(obj.Pop())
    fmt.Println(obj.Pop())
    fmt.Println(obj.Empty())
}
