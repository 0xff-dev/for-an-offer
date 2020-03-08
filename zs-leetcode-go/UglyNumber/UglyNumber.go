package UglyNumber

func isUgly(num int) bool {
    for num % 2 == 0 {
        num /= 2
    }
    for num % 3 == 0 {
        num /= 3
    }
    for num %5 == 0 {
        num /= 5
    }
    return num == 1
}
