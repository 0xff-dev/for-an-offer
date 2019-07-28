package EvaluateReversePolishNotation

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	input := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(input))
	input = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(input))
}
