package EvaluateDivision


func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    res := make([]float64, 0)
    length := len(equations)
    charMap := make(map[string]float64)
    for index := 0; index < length; index++ {
        if val, ok := charMap[equations[index][0]]; ok {
            charMap[equations[index][1]] = val / values[index] * 1.0
            continue
        } else if val, ok := charMap[equations[index][1]]; ok {
            charMap[equations[index][0]] = val * values[index]
        } else {
            charMap[equations[index][1]] = 1.0
            charMap[equations[index][0]] = values[index]
        }
    }
    for _, item := range queries {
        dividend, ok0 := charMap[item[0]]
        divisor, ok1 := charMap[item[1]]
        if !ok0 || !ok1 {
            res = append(res, -1.0)
            continue
        }
        res = append(res, dividend/divisor*1.0)
    }
    return res
}
