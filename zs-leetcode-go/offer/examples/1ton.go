package examples

import (
	"fmt"
	"strconv"
)

func Print1ton(num int) {
	str := fmt.Sprintf("%d", num)
	fmt.Println(calculate(str))
}

func calculate(num string) int {
	if num[0] < '0' || num[0] > '9' || len(num) == 0 {
		return 0
	}
	first := num[0] - '0'
	length := len(num)
	if length == 1 && first == 0 {
		return 0
	}
	if length == 1 && first > 1 {
		return 1
	}
	// 最高位的1个数
	var highBitNum int
	if first > 1 {
		highBitNum = powBase10(length - 1)
	} else {
		res, _ := strconv.Atoi(num[1:])
		highBitNum = res + 1
	}
	// 计算部分
	notHighBit := int(first) * powBase10(length-2) * (length - 1)
	lowBit := calculate(num[1:])
	return highBitNum + notHighBit + lowBit
}

func powBase10(exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= 10
	}
	return res
}
