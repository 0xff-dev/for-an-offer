package LongestWordInDictionary

import "testing"

func TestLongestWord(t *testing.T) {
	input := []string{"w", "wo", "wor", "worl", "world"}
	if res := longestWord(input); res != "world" {
		t.Fatalf("expect %s get %s", "world", res)
	}
	input = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	if res := longestWord(input); res != "apple" {
		t.Fatalf("expect %s get %s", "apple", res)
	}
}
