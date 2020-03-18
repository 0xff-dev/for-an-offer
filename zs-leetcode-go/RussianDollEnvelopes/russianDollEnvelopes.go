package RussianDollEnvelopes

import "sort"
func maxEnvelopes(envelopes [][]int) int {
    if len(envelopes) == 0 { return 0 }
    sort.Sort(sortTyep(envelopes))
    maxEnves := 0
    res := make([]int, len(envelopes))
    for index := 1; index < len(envelopes); index++ {
        for inner := 0; inner < index; inner ++ {
            if less(envelopes[inner], envelopes[index]){
                res[index] = max(res[index], res[inner]+1)
            }
        }
        maxEnves = max(maxEnves, res[index])
    }
    return maxEnves+1
}

type sortTyep [][]int
func (self sortTyep) Len() int {
    return len(self)
}
func (self sortTyep) Less(i, j int) bool {
    if self[i][0] == self[j][0] {
        return self[i][1] < self[j][1]
    }
    return self[i][0] < self[j][0]
}
func (self sortTyep) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}

func less(a, b []int) bool {
    if a[0] < b[0] && a[1] < b[1] { return true }
    return false
}
func max(a, b int) int {
    if a < b { return b }
    return a
}
