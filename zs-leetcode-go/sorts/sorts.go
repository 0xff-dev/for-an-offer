/*
	实现常用的排序算法
	1. 快速排序
	2. 归并排序
	3. 堆排序
	4. 冒泡排序
	5. 选择排序
	6. 插入排序
	7. 睡眠排序
*/
package sorts

import (
	"fmt"
	"sync"
	"time"
)

// 寻找分隔index
func Paritation(nums []int, left, right int) int {
	i, j := left, left+1
	compareNum := nums[left]
	for ; j <= right; j++ {
		if nums[j] <= compareNum {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

// 快速排序两种实现形式
func QuickSort(nums []int, left, right int) {
	//if left < right {
	//	i, j := left, right
	//	compareNum := nums[left]
	//	for i < j {
	//		for j > i && nums[j] > compareNum {
	//			j--
	//		}
	//		for i < j && nums[i] < compareNum {
	//			i++
	//		}
	//		if i < j {
	//			nums[i], nums[j] = nums[j], nums[i]
	//		}
	//	}
	//	QuickSort(nums, left, i)
	//	QuickSort(nums, i+1, right)
	//}
	if left < right {
		index := Paritation(nums, left, right)
		QuickSort(nums, left, index)
		QuickSort(nums, index+1, right)
	}
}

// 归并排序归并部分
func merge(nums []int, left, right int) {
	mid := (left + right) / 2
	i, j := left, mid+1
	index := 0
	newArray := make([]int, right-left+1)
	for ; i <= mid && j <= right; index++ {
		if nums[i] < nums[j] {
			newArray[index] = nums[i]
			i++
		} else {
			newArray[index] = nums[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		newArray[index] = nums[i]
		index++
	}
	for ; j <= right; j++ {
		newArray[index] = nums[j]
		index++
	}
	for k := 0; k < index; k++ {
		nums[left+k] = newArray[k]
	}
}

// 归并排序
func MergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(nums, left, mid)
		MergeSort(nums, mid+1, right)
		merge(nums, left, right)
	}
}

// 堆排序调整部分
func maintain(nums []int, index, heapSize int) {
	left, right := index*2+1, (index+1)*2
	aimIndex := index
	if left <= heapSize && nums[left] > nums[aimIndex] {
		aimIndex = left
	}
	if right <= heapSize && nums[right] > nums[aimIndex] {
		aimIndex = right
	}
	if aimIndex != index {
		nums[index], nums[aimIndex] = nums[aimIndex], nums[index]
		maintain(nums, aimIndex, heapSize)
	}
}

// 构建堆
func buidlHeap(nums []int) {
	length := len(nums) - 1
	for i := length / 2; i >= 0; i-- {
		maintain(nums, i, length)
	}
}

// 堆排序
func HeapSort(nums []int) {
	buidlHeap(nums)
	length := len(nums) - 1
	for length > 0 {
		nums[0], nums[length] = nums[length], nums[0]
		length--
		maintain(nums, 0, length)
	}
}

// 咕噜咕噜排序
func BubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 选择排序
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// 插入排序
func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				tmp := nums[i]
				for k := i; k > j; k-- {
					nums[k] = nums[k-1]
				}
				nums[j] = tmp
				break
			}
		}
	}
}

// 睡眠排序
func SleepSort(nums []int) {
	var wg sync.WaitGroup
	wg.Add(len(nums))
	for _, item := range nums {
		go func(item int) {
			defer wg.Done()
			time.Sleep(time.Duration(item) * time.Second)
			fmt.Println(item)
		}(item)
	}
	wg.Wait()
}
