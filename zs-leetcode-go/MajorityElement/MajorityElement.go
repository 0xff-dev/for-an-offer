package MajorityElement


func majorityElement(nums []int) int {
    times := 1
    aim := nums[0]
    for _, item := range nums[1:] {
        if item == aim {
            times++
        } else {
            times--
        }
        if times == 0 {
            aim = item
            times = 1
        }
    }
    return aim
}
