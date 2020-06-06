package HappyNumber

//不知道这道题的规律，只能硬做了。。
func isHappy(n int) bool {
	aim := n
	checkMap := map[int]struct{}{}
	for {
		tmp := calculate(aim)
		if judge(tmp){
			return true
		}
		if tmp == n {
			return false
		}
		if _, ok := checkMap[tmp]; ok {
			return false
		}
		checkMap[tmp] = struct{}{}
		aim = tmp
	}
}


func calculate(n int) int {
	res := 0
	for n > 0 {
		tmp := n%10
		res += tmp*tmp
		n/=10
	}
	return res
}

func judge(n int) bool {
	res := 1
	for {
		if res == n {
			return true
		}
		if res > n {
			break
		}
		res *= 10
	}
	return false
}