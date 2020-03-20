package TeemoAttacking

import (
    "fmt"
    "testing"
)

func TestFindPoisoneDuration(t *testing.T) {
    case1 := []int{1, 4}
    case2 := []int{1, 2}
    fmt.Println(findPoisonedDuration(case1, 2))
    fmt.Println(findPoisonedDuration(case2, 2))
    cast3 := []int{1}
    fmt.Println(findPoisonedDuration(cast3, 1))
}
