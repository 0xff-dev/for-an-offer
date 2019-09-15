package ReverseBits

// 位运算
// 4294967295
func reverseBits(num uint32) uint32 {
	constNums := []uint32{0xAAAAAAAA, 0x55555555, 0xCCCCCCCC, 0x33333333, 0xF0F0F0F0, 0x0F0F0F0F, 0xFF00FF00, 0x00FF00FF, 0xFFFF0000, 0x0000FFFF}
	var index, power uint32
	for index, power = 0, 0; index <= 10; index, power = index+2, power+1 {
		num = ((num & constNums[index]) >> (1 << power)) | ((num & constNums[index+1]) << (1 << power))
	}
	return num
}
