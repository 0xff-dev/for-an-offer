package Numberof1Bits

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += 1
		num = num & (num - 1)
	}
	return res
}
