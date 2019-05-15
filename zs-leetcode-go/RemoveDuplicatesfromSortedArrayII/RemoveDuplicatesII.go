//  O(1)的额外空间是在原数组修改， 那么我们就要对数组进行移动..
package RemoveDuplicatesfromSortedArrayII

func removeDuplicates(nums []int) int {
	length := len(nums)
	for index := 0; index < length-1; index++ {
		cnt := 1
		for j := index + 1; j < length; j++ {
			if nums[j] == nums[index] {
				cnt += 1
			} else {
				break
			}
		}
		if cnt > 2 {
			steps := cnt - 2
			start := index + 2
			for k := index + cnt; k < length; k, start = k+1, start+1 {
				nums[start] = nums[k]
			}
			length -= steps
			index++
		}
	}
	nums = nums[:length]
	return length
}
