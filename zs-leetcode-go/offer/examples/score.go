// n个骰子， 组成的点数
package examples

import "fmt"

var maxPoint = 6

func BeautifulSolve(n int) {
	if n == 0 {
		return
	}
	points := make([][]int, 2)
	for i := 0; i < 2; i++ {
		points[i] = make([]int, maxPoint*n+1)
	}
	flag := 0
	for i := 1; i <= 6; i++ {
		points[flag][i] = 1
	}
	for k := 2; k <= n; k++ {
		for index := 0; index < k; index ++ {
			points[1-flag][index] = 0
		}
		for index := k; index <= k*maxPoint; index ++ {
			points[1-flag][index] = 0
			for j := 1; j <= index && j <= maxPoint; j++ {
				points[1-flag][index] += points[flag][index-j]
			}
		}
		flag = 1- flag
	}
	for i := n; i <= n*maxPoint; i++ {
		fmt.Printf("point:= %d; get:= %d\n", i, points[flag][i])
	}
}