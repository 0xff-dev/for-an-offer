package InsertInterval

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	testDatas := []Interval{
		{1, 3},
		{6, 9},
	}
	fmt.Println(insert(testDatas, Interval{2, 5}))
	testDatas = []Interval{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	fmt.Println(insert(testDatas, Interval{4, 8}))

	testDatas = []Interval{}
	fmt.Println(insert(testDatas, Interval{5, 7}))

	testDatas = []Interval{{1, 5}}
	fmt.Println(insert(testDatas, Interval{4, 8}))

	testDatas = []Interval{
		{1, 5},
	}
	fmt.Println(insert(testDatas, Interval{6, 8}))
}
