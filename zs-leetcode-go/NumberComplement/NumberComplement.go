package NumberComplement

func findComplement(num int) int {
	start := 2
	for ;start <= num; start*=2 {}
	if num == start-1 { return 0}
	return start-1-num
}
