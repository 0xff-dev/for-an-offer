package IntersectionOfTwoArraysII

func intersect(nums1 []int, nums2 []int) []int {
    one, two := make(map[int][]int), make(map[int][]int) 
    for _, item := range nums1 {
        if _, ok := one[item]; !ok {
            one[item] = []int{}
        }
        one[item] = append(one[item], item)
    }

    for _, item := range nums2 {
        if _, ok := two[item]; !ok {
            two[item] = []int{}
        }
        two[item] = append(two[item], item)
    }

    res := make([]int, 0)
    for k, v := range one {
        if val, ok := two[k]; ok {
            if len(val) < len(v) {
                res = append(res, val...)
            } else {
                res = append(res, v...)
            }
        }
    }
    return res
}
