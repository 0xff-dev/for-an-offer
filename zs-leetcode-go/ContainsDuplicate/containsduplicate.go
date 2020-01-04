package ContainsDuplicate

func containsDuplicate(nums []int) bool {
    checkMap := make(map[int]bool)
    for _, item := range nums {
        if _, ok := checkMap[item]; ok {
            return true
        }
        checkMap[item] = true
    }
    return false
}
