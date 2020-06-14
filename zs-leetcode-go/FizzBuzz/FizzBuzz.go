package FizzBuzz

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, 0)
	for index := 1; index <=n; index ++ {
		var str string
		if index % 3 == 0 && index % 5 == 0 {
			str = "FizzBuzz"
		} else if index % 3 == 0 {
			str = "Fizz"
		} else if index % 5 == 0 {
			str = "Buzz"
		} else {
			str = fmt.Sprintf("%d", index)
		}
		res = append(res, str)
	}
	return res
}