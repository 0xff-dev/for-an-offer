package BinarySearch

import "testing"

func TestSearch(t *testing.T) {
	input := []int{-1, 0, 3, 5, 9, 12}
	if search(input, 9) == -1 {
		t.Fatal("not found")
	}
	input = []int{}
	if search(input, 10) != -1 {
		t.Fatal("found error")
	}
	input = []int{5}
	if search(input, 5) == -1 {
		t.Fatal("found error")
	}
}
