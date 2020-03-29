package ThridMaximumNumber
func thirdMax(nums []int) int {
    length := len(nums)
    res := []int{}
    for index := 0; index < length; index ++ {
        // adjust array
        insert(&res, nums[index])
    }
    if len(res) < 3 { return res[0] }
    return res[2]
}

// len(data) <= 3
func insert(data *[]int, val int) {
    index, length, end := 0, len(*data), 2
    if length < 3 { end = length }
    inArray := false
    for ; index < len(*data) && index < 3; index ++ {
        if (*data)[index] == val {
            inArray = true
            break
        }
        if (*data)[index] < val {
            break
        }
    }
    if index == 3 { return  }
    if inArray { return }
    if index == len(*data) {
        *data = append(*data, val)
        return 
    }
    *data = append((*data)[:index], append([]int{val}, (*data)[index:end]...)...)
}
