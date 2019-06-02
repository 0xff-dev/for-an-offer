package PascalTriangle


func addZero(list []int) []int {
	tmp := make([]int, 0)
	tmp = append(tmp, 0)
	tmp = append(tmp, list...)
	tmp = append(tmp, 0)
	return tmp
}

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	tmpList := []int{0, 1, 0}
	for row := 1; row < numRows; row++ {
		tmp := make([]int, 0)
		for index := 0; index < len(tmpList)-1; index++ {
			tmp = append(tmp, tmpList[index]+tmpList[index+1])
		}
		res = append(res, tmp)
		tmpList = addZero(tmp)
	}
	return res
}
