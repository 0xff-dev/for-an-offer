package NumberOfDigitOne

import (
	"fmt"
	"strconv"
)

func countDigitOne(n int) int {
	if n < 0 {
		return 0
	}
	return aux(fmt.Sprintf("%d", n))
}

func aux(num string) int {
	numLen := len(num)
	if numLen == 0 {
		return 0
	}
	first := num[0]
	if first == '0' && numLen == 1 {
		return 0
	}
	if first > '0' && numLen == 1 {
		return 1
	}
	hight := 0
	if first > '1' {
		hight = powerBase10(numLen - 1)
	} else if first == '1' {
		hight, _ = strconv.Atoi(num[1:])
		hight += 1
	}
	others := int(first-'0') * (numLen - 1) * powerBase10(numLen-2)
	recursiveRes := aux(num[1:])
	return hight + others + recursiveRes
}

func powerBase10(x int) int {
	res := 1
	for index := 0; index < x; index++ {
		res *= 10
	}
	return res
}
