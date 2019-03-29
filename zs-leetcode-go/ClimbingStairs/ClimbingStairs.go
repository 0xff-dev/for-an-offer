package ClimbingStairs


func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	one, two := 1, 2
	for i := 3; i<=n; i++ {
		tmp := one+two
		one = two
		two = tmp
	}
	return two
}