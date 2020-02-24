package ValidPerfectSquare

func isPerfectSquare(num int) bool {
    if num < 2 { return true }
    left, right := 1, num
    for right-left > 1 {
        mid := left + (right-left)/2
        // mid * mid 越界问题？
        data := mid * mid
        if data == num {
            return true
        }
        if data < num {
            left = mid
        } else {
            right = mid
        }
    }
    return false
}
