package bitcharacters

import "testing"

func TestIsOnBitCharacter(t *testing.T) {
    input := []int{1,1,0}
    if isOneBitCharacter(input) != true {
        t.Fatal("except true")
    }
    input = []int{1,1,1,0}
    if isOneBitCharacter(input) != false {
        t.Fatal("except false")
    }
    input = []int{1}
    if isOneBitCharacter(input) != false {
        t.Fatal("except false")
    }
    input = []int{0}
    if isOneBitCharacter(input) != true {
        t.Fatal("except true")
    }
    if isOneBitCharacter([]int{}) != false {
        t.Fatal("empty array except false")
    }
}
