package LongestConsecutiveSequence

// 空间换时间的问题 像是dp的问题呢
func longestConsecutive(nums []int) int {
	resMap := make(map[int]bool)
	for _, item := range nums {
		resMap[item] = true
	}
	cnt := 0
	for _, item := range nums {
		tmpCnt, now := 1, item+1
		for ; ; now, tmpCnt = now+1, tmpCnt+1 {
			if _, ok := resMap[now]; !ok {
				break
			}
		}
		if tmpCnt > cnt {
			cnt = tmpCnt
		}
	}
	return cnt
}
