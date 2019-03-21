// 两种方法解决
// 1. 数组计算， 打印
// 2. 全排列
package examples

import "fmt"

// 判断数组是不是都是9
func compare(nums []int) bool {
	for _, i := range nums {
		if i != 9 {
			return true
		}
	}
	return false
}

func printArray(nums []int) {
	firstZero := true
	for _, i := range  nums {
		if i == 0 && firstZero {
			continue
		} else {
			fmt.Print(i)
			firstZero = false
		}
	}
	fmt.Print(" ")
}

// 数组加法
func addArray(nums []int) {
	index := len(nums)-1
	cf := 0
	tmp := nums[index]+1
	if tmp >= 10 {
		nums[index] = tmp % 10
		cf = tmp / 10
		for i := index-1; i >= 0; i-- {
			tmp = nums[i] + cf
			if tmp >= 10 {
				nums[i] = tmp % 10
				cf = tmp / 10
			} else {
				nums[i] = tmp
				break
			}
		}
	} else {
		nums[index] = tmp
	}
}

func OneToNBit(integer int) {
	res := make([]int, integer)
	for compare(res) {
		addArray(res)
		printArray(res)
	}
	fmt.Println()
}


func SortAllInteger(integer int) {
	res := make([]int, integer)
	for i := 0; i < 10; i++ {
		res[0] = i
		auxCalculate(res, 0, integer)
	}
}

func auxCalculate(nums []int, start, integer int) {
	if start == integer-1 {
		printArray(nums)
		return
	}
	for i := 0; i < 10; i++ {
		nums[start+1] = i
		auxCalculate(nums, start+1, integer)
	}
}