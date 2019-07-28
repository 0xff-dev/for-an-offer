package BasicCalculator

import (
	"strconv"
)


func isNums(char byte) bool {
	return char >= '0' && char <= '9'
}

func cal(opn1, opn2 int, opt, pre byte) int {
	// - - opn1+opn2
	// - + opn2-opn1
	// + + opn1+opn2
	// + - opn2-opn1
	// ( + opn1+opn2
	// ( - opn2-opn1
	if pre == opt || pre == '(' && opt =='+' {
		return opn1+opn2
	}
	return opn2-opn1
}


func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	pre := byte('(')
	charStack := make([]byte, 0)
	charSize := 0
	numberStack := make([]int, 0)
	numberSize := 0
	index := 0
	// - 的计算高于 +
	for index<len(s) {
		if s[index] != ' ' {
			if s[index] == '(' {
				charSize += 1
				charStack = append(charStack, s[index])
			} else if s[index] == ')' {
				for ;charStack[charSize-1] != '('; charSize-- {
					pre = '('
					if charSize-2 >= 0 {
						pre = charStack[charSize-2]
					}
					opn1, opn2 := numberStack[numberSize-1], numberStack[numberSize-2]
					num := cal(opn1, opn2, charStack[charSize-1], pre)
					numberStack = numberStack[:numberSize-2]
					numberStack = append(numberStack, num)
					numberSize -= 1
				}
				charSize --
				charStack = charStack[:charSize]
			} else {
				if !isNums(s[index]) {
					charSize++
					charStack = append(charStack, s[index])
				} else {
					start := index
					index += 1
					for ; index < len(s) && isNums(s[index]); index++ {}
					num, _ := strconv.Atoi(s[start:index])
					numberSize++
					numberStack = append(numberStack, num)
					index -= 1
				}
			}
		}
		index++
	}
	for ;charSize != 0; charSize-- {
		pre = '('
		if charSize-2 >= 0 {
			pre = charStack[charSize-2]
		}
		opn1, opn2 := numberStack[numberSize-1], numberStack[numberSize-2]
		num := cal(opn1, opn2, charStack[charSize-1], pre)
		numberStack = numberStack[:numberSize-2]
		numberStack = append(numberStack, num)
		numberSize -= 1
	}
	return numberStack[0]
}
