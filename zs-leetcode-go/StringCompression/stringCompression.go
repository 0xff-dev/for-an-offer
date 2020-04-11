package StringCompression

func compress(chars []byte) int {
	length, index, pos := len(chars), 1, 0
	if length == 0 {
		return 0
	}
	var count int = 0
	for index < len(chars) {
		if chars[index] == chars[pos] {
			count, index = count+1, index+1
			continue
		}
		if count != 0 {
			newBytes := converIntToBytes(count + 1)
			chars = append(chars[:pos+1], append(newBytes, chars[index:]...)...)
			index = index + len(newBytes) - count
			count = 0
		}
		pos = index
		index++
	}
	if pos != len(chars)-1 {
		chars = append(chars[:pos+1], converIntToBytes(count+1)...)
	}
	return len(chars)
}

// convertIntToBytes convert a unit8 to bytes array
func converIntToBytes(num int) []byte {
	res := make([]byte, 0)
	for num >= 10 {
		tmp := num % 10
		num /= 10
		res = append(res, uint8(tmp)+'0')
	}
	res = append(res, uint8(num)+'0')
	for start, end := 0, len(res)-1; start < end; start, end = start+1, end-1 {
		res[start], res[end] = res[end], res[start]
	}
	return res
}
