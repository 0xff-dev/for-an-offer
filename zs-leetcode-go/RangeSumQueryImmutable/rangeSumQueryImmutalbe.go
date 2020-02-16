package RangeSumQueryImmutable


type NumArray struct {
    Sums []int
    Length int
}

func Constructor(nums []int) NumArray {
    length := len(nums)
    sums := make([]int, length)
    if length == 0 {
        return NumArray{
            Sums: sums,
            Length: length,
        }
    }
    sums[0] = nums[0]
    for index := 1; index < length; index ++ {
        sums[index] = sums[index-1]+nums[index]
    }
    return NumArray{
        Sums: sums,
        Length: length,
    }
}

func (this *NumArray) SumRange(i int, j int) int {
    if this.Length == 0 { return 0 } 
    if i == 0 { return this.Sums[j] }
    return this.Sums[j] - this.Sums[i-1]
}

