package BinaryWatch

import "fmt"

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	hourMap := map[int][]int{
		0: []int{0},
		1: []int{1, 2, 4, 8},
		2: []int{3, 5, 6, 9, 10},
		3: []int{7, 11},
	}
	res := []string{}
	oneMap := map[int][]int{}
	for i := 0; i < 60; i++ {
		count := cntOfOne(i)
		if _, ok := oneMap[count]; !ok {
			oneMap[count] = []int{}
		}
		oneMap[count] = append(oneMap[count], i)
	}
	for length := 0; length <= num && length < 4; length++ {
		merge(hourMap[length], oneMap[num-length], &res)
	}
	return res
}

func cntOfOne(num int) int {
	cnt := 0
	for num > 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}

func merge(hour, minute []int, res *[]string) {
	formate := "%d:%d"
	ltTenFormate := "%d:0%d"
	var str string
	for _, h := range hour {
		for _, m := range minute {
			str = fmt.Sprintf(formate, h, m)
			if m < 10 {
				str = fmt.Sprintf(ltTenFormate, h, m)
			}
			*res = append(*res, str)
		}
	}
}
