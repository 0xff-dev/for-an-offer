package SumOfTwoIntegers


func getSum(a int, b int) int {
    sum := a ^ b
    carry := (a & b) << 1
    for b != 0 {
        a, b = sum, carry
        sum = a ^ b
        carry = (a & b) << 1
    }
    return a
}
