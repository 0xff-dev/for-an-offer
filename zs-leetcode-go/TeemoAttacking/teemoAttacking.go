package TeemoAttacking


func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 0 { return 0 }
    attackTime := 0
    for index := 1; index < len(timeSeries); index ++ {
        tmpDuration := timeSeries[index] - timeSeries[index-1]
        if tmpDuration > duration {
            tmpDuration = duration
        }
        attackTime += tmpDuration
    }
    attackTime += duration
    return attackTime
}
