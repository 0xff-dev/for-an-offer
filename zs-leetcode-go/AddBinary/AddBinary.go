package AddBinary

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string{
	aLen := len(a)
	bLen := len(b)
	length := 0
	if aLen > bLen {
		length = aLen
	} else {
		length = bLen
	}
	aIndex, bIndex := aLen-1, bLen-1
	index := length-1
	tmp := -1
	res := make([]string, length)
	cf := 0
	// 1011+11 = 1110
	for ;aIndex >= 0 && bIndex >= 0; aIndex, bIndex = aIndex-1, bIndex-1 {
		tmp = int(a[aIndex]-'0') + int(b[bIndex]-'0') + cf
		cf = tmp / 2
		if tmp >= 2 {
			tmp -= 2
		}
		res[index] = fmt.Sprintf("%d", tmp)
		index--
	}
	for ;aIndex >= 0; aIndex-- {
		tmp = int(a[aIndex]-'0') + cf
		cf = tmp / 2
		if tmp >= 2 {
			tmp -= 2
		}
		res[index] = fmt.Sprintf("%d", tmp)
		index--
	}
	for ;bIndex >= 0; bIndex-- {
		tmp = int(b[bIndex]-'0') + cf
		cf = tmp / 2
		if tmp >= 2 {
			tmp -= 2
		}
		res[index] = fmt.Sprintf("%d", tmp)
		index--
	}
	if cf == 0 {
		return strings.Join(res, "")
	} else {
		return fmt.Sprintf("%d", cf) + strings.Join(res, "")
	}
}
