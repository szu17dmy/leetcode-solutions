package climbing_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	p, pp := 1, 2
	for i := 3; i < n; i++ {
		p, pp = pp, p+pp
	}
	return p + pp
}
