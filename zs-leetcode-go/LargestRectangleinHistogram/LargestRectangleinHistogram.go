package LargestRectangleinHistogram

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxArea := 0
	for outIndex := 0; outIndex < len(heights); outIndex++ {
		cnt := 1
		for index := outIndex+1; index < len(heights); index++ {
			if heights[index] >= heights[outIndex] {
				cnt ++
			} else {
				break
			}
		}
		for index := outIndex-1; index >= 0; index-- {
			if heights[index] >= heights[outIndex] {
				cnt ++
			} else {
				break
			}
		}
		if cnt*heights[outIndex]> maxArea {
			maxArea = cnt * heights[outIndex]
		}
	}
	return maxArea
}
