package IntersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	nMap1, nMap2 := make(map[int]int), make(map[int]int)
	for _, item := range nums1 {
		nMap1[item]++
	}
	for _, item := range nums2 {
		nMap2[item]++
	}
	res := make([]int, 0)
	for k := range nMap1 {
		if _, ok := nMap2[k]; ok {
			res = append(res, k)
		}
	}
	return res
}
