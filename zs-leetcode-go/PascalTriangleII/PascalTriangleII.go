package PascalTriangleII


func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := make([]int, rowIndex+1)
	res[0] = 1
	for cnt := 1; cnt <= rowIndex; cnt++ {
		res[cnt] = 1
		for start := cnt-1; start > 0; start--{
			res[start] += res[start-1]
		}

	}
	return res
}
