package ScoreOfParentheses

import "testing"

func TestScoreOfParentheses(t *testing.T) {
	input := "()"
	if scoreOfParentheses(input) != 1 {
		t.Fatal(input, "except 1")
	}
	input = "(())"
	if scoreOfParentheses(input) != 2 {
		t.Fatal(input, "except 2")
	}
	input = "()()"
	if scoreOfParentheses(input) != 2 {
		t.Fatal(input, "except 2")
	}
	input = "(()(()))"
	if scoreOfParentheses(input) != 6 {
		t.Fatal(input, "except 6")
	}
}
