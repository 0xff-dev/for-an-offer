package ExcelSheetColumnNumber

func titleToNumber(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	cf := 1
	res := 0
	for index := length-1; index >= 0; index-- {
		res += (int(s[index]-'A')+1)*cf
		cf *= 26
	}
	return res
}
