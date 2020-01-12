package ImplementStackUsingQueues

import (
    "fmt"
    "testing"
)
func TestImplementStack(t *testing.T) {
    stack := Constructor()
    stack.Push(1)
    stack.Push(2)
    fmt.Println(stack.Top())
    fmt.Println(stack.Pop())
    fmt.Println(stack.Empty())
}
