// 大数的乘法， 每个至多110位
package MultiplyStrings

import (
	"fmt"
	"strings"
)

func reverseString(nums string) []string {
	bytes := strings.Split(nums, "")
	i, j := 0, len(nums)-1
	for ;i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
func multiply(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2)+1)
	arrNum1, arrNum2 := reverseString(num1), reverseString(num2)
	cf := 0
	for out, outItem := range arrNum1 {
		cf = 0
		for in, inItem := range arrNum2 {
			tmp := int(outItem[0]-'0') * int(inItem[0]-'0') + cf
			res[in+out] += tmp
			cf = res[in+out] / 10
			res[in+out] %= 10
		}
		res[out+len(arrNum2)] = cf
	}
	returnStr, first := "", true
	for i := len(res)-1; i >= 0; i-- {
		if first && res[i] == 0 {
			continue
		} else {
			first = false
			returnStr += fmt.Sprintf("%d", res[i])
		}
	}
	if first {return "0"}
	return returnStr
}