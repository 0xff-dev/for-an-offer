package RansomNote

import "testing"

func TestCanConstruct(t *testing.T) {
	input, magazine := "a", "b"
	if canConstruct(input, magazine) != false {
		t.Fatalf("except false")
	}
	input, magazine = "aa", "ab"
	if canConstruct(input, magazine) != false {
		t.Fatalf("except false")
	}
	input, magazine = "aa", "aab"
	if canConstruct(input, magazine) != true {
		t.Fatalf("except true")
	}
}
