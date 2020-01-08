package CountPrimes

func countPrimes(n int) int {
    if n < 2 {
        return 0
    }
    length := 0
    mark := make([]bool, n)
    for index := 2; index < n; index++{
        for inner := index+index; inner < n; inner += index {
            mark[inner] = true
        }
    } 
    for _, item := range mark[2:] {
        if !item { length++ }
    }
    return length
}
