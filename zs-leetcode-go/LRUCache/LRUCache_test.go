package LRUCache

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	root := Constructor(2)
	root.Put(1, 1)
	root.Put(2, 2)
	fmt.Println(root.Get(1))
	root.Put(3, 3)
	fmt.Println(root.Get(2))
	root.Put(4, 4)
	fmt.Println(root.Get(1))
	fmt.Println(root.Get(3))
	fmt.Println(root.Get(4))

	root = Constructor(2)
	root.Put(2, 1)
	root.Put(1, 1)
	root.Put(2, 3)
	root.Put(4, 1)
	fmt.Println(root.Get(1))
	fmt.Println(root.Get(2))

	root = Constructor(2)
	root.Put(2, 6)
	fmt.Println(root.Get(1))
	root.Put(1, 5)
	root.Put(1, 2)
	fmt.Println(root.Get(1))
	fmt.Println(root.Get(2))
}
