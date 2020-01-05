package ContainsDuplicateII

func containsNearbyDuplicate(nums []int, k int) bool {
    checkMap := make(map[int]int) // value 代表index的最小值
    for index, item := range nums {
        if val, ok := checkMap[item]; !ok {
            checkMap[item] = index 
        } else {
            if index - val <= k {
                return true
            }
            checkMap[item] = index
        }
    }
    return false
}
