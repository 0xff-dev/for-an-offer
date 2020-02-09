package SummaryRanges

import "fmt"
func summaryRanges(nums []int) []string {
    length := len(nums)
    if length == 0 { return nil }
    if length == 1 { return []string{fmt.Sprintf("%d", nums[0])} }
    res, start := make([]string, 0), 0
    for index := 1; index < length; index ++ {
        if nums[index]-nums[index-1] !=1 {
            var str string
            if index - 1 == start {
                str = fmt.Sprintf("%d", nums[start])
            } else {
                str = fmt.Sprintf("%d->%d", nums[start], nums[index-1])
            }
            res = append(res, str)
            start = index
        }
    }
    if start == length -1 {
        res = append(res, fmt.Sprintf("%d", nums[start]))
    } else {
        res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[length-1]))
    }
    return res
}
