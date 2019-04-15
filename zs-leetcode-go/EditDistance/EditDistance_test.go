package EditDistance

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	word1, word2 := "horse", "ros"
	fmt.Println(minDistance(word1, word2))
	word1 = "intention"
	word2 = "execution"
	fmt.Println(minDistance(word1, word2))
}
