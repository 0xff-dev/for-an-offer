package sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testDatas := []int{10, 4, 34, 90, 23, 98}
	QuickSort(testDatas, 0, len(testDatas)-1)
	fmt.Println("QuickSort: ", testDatas)
	testDatas = []int{45, 54, 12, 34, 90}
	MergeSort(testDatas, 0, len(testDatas)-1)
	fmt.Println("MergeSort: ", testDatas)
	testDatas = []int{100, 45, 122, 56, 8, 12}
	HeapSort(testDatas)
	fmt.Println("HeapSort: ", testDatas)
	testDatas = []int{99, 44, 123, 456, 25, 98}
	HeapSort(testDatas)
	fmt.Println("HeapSort: ", testDatas)
	testDatas = []int{99, 44, 123, 456, 25, 98}
	BubbleSort(testDatas)
	fmt.Println("Bubble: ", testDatas)
	testDatas = []int{99, 44, 123, 456, 25, 98}
	SelectSort(testDatas)
	fmt.Println("SelectSort: ", testDatas)
	testDatas = []int{99, 44, 123, 456, 25, 98}
	InsertSort(testDatas)
	fmt.Println("InsertSort: ", testDatas)
	testDatas = []int{12, 8, 3, 6}
	SleepSort(testDatas)
}
