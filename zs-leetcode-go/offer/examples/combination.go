/*
	 实现排列组合， 全排列算法
*/
package examples


import (
	"fmt"
	"strings"
)

func Combination(str string) {
	stack := make([]byte, 0)
	for length := 1; length <= len(str); length ++ {
		combination(str, 0, length, stack)
	}
}

func combination(str string, index, length int, stack []byte) {
	if length == 0 {
		for _, item := range stack {
			fmt.Printf("%c ", item)
		}
		fmt.Println()
		return
	}
	// 选择与补选的情况
	if index >= len(str) {
		return
	}
	stack = append(stack, byte(str[index]))
	combination(str, index+1, length-1, stack)
	stack = stack[:len(stack)-1]
	combination(str, index+1, length, stack)
}

func AllSort(str string) {
	strs := strings.Split(str, "")
	search(strs, 0, len(str))
}

func search(strs []string, index, length int) {
	if index == length {
		fmt.Println(strs)
		return
	}
	for i := index; i <length; i++ {
		strs[i], strs[index] = strs[index], strs[i]
		search(strs, index+1, length)
		strs[i], strs[index] = strs[index], strs[i]
	}
}