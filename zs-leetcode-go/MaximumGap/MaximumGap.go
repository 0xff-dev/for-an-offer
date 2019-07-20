package MaximumGap

// 1. 排序，计算
// 2. 桶排序思想
// 3. 解法给的第三种方法，学习了
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func getMaxAndMin(nums []int) (_max, _min int) {
	_max, _min = nums[0], nums[0]
	for _, item := range nums[1:] {
		_max = max(item, _max)
		_min = min(item, _min)
	}
	return
}

type Bucket struct {
	Used bool
	Min, Max int
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	_max, _min := getMaxAndMin(nums)
	size := max(1, (_max-_min)/(len(nums)-1))
	buckets := make([]Bucket, 0)
	for i := 0; i < (_max-_min)/size+1; i++ {
		buckets = append(buckets, Bucket{false, 100000000, -1})
	}
	for _, item := range nums {
		index := (item-_min)/size
		buckets[index].Used = true
		buckets[index].Min = min(item, buckets[index].Min)
		buckets[index].Max = max(item, buckets[index].Max)
	}
	pre, maxGap := _min, 0
	for _, item := range buckets {
		if !item.Used {
			continue
		}
		maxGap = max(maxGap, item.Min-pre)
		pre = item.Max
	}
	return maxGap
}
