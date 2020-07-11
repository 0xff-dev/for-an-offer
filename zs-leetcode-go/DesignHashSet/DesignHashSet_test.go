package DesignHashSet

import "testing"

func TestDesignHashSet(t *testing.T) {
	hash := Constructor()
	hash.Add(1)
	hash.Add(2)
	if hash.Contains(1) != true {
		t.Fatal("except true")
	}
	if hash.Contains(3) != false {
		t.Fatal("except false")
	}
	hash.Add(2)
	if hash.Contains(2) != true {
		t.Fatal("except true")
	}
	hash.Remove(2)
	if hash.Contains(2) != false {
		t.Fatal("except false")
	}
}
