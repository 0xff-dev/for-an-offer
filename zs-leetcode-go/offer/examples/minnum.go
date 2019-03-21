package examples

import (
	"fmt"
	"sort"
	"strings"
)

type compareString []string

func (ci compareString) Len() int {
	return len(ci)
}

func (ci compareString) Swap(i, j int) {
	ci[i], ci[j] = ci[j], ci[i]
}

func (ci compareString) Less(i, j int) bool {
	return ci[i]+ci[j] < ci[j]+ci[i]
}

func Minnum(nums []int) {
	strs := make([]string, 0)
	for _, item := range nums {
		strs = append(strs, fmt.Sprintf("%d", item))
	}
	sort.Sort(compareString(strs))
	fmt.Println(strings.Join(strs, ""))
}
