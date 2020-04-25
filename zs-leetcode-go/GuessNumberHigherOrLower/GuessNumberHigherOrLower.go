package GuessNumberHigherOrLower

import (
	"math/rand"
	"time"
)

func guessNumber(n int) int {
	left, right := 1, n
	var mid int
	for left <= right {
		mid = (right - left)/2 + left
		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// this function just for test,
func guess(num int) int {
	result := []int{-1, 0, 1}
	return result[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3)]
}