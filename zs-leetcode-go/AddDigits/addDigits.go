package AddDigits


func addDigits(num int) int {
    for num > 9 {
        num = num%10+num/10
    }
    return num
}
