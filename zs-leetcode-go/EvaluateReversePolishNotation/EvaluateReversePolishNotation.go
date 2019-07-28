package EvaluateReversePolishNotation

import (
	"strconv"
)

func isOperator(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "/" {
		return true
	}
	return false
}


func calculate(opn1, opn2 int, opt string) int{
	switch opt {
	case "+":
		return opn1+opn2
	case "-":
		return opn2-opn1
	case "*":
		return opn1*opn2
	case "/":
		return opn2/opn1
	default:
		return 0
	}
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	size := 0
	for _, item := range tokens {
		if isOperator(item) {
			opn1, opn2 := stack[size-1], stack[size-2]
			stack = stack[:size-2]
			stack = append(stack, calculate(opn1, opn2, item))
			size -= 1
		} else {
			num, _ := strconv.Atoi(item)
			stack = append(stack, num)
			size ++
		}
	}
	return stack[0]
}