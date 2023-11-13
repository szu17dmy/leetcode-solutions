package rotate_array

func rotate(nums []int, k int) {
	k %= len(nums)
	for b := 0; b < gcd(k, len(nums)); b++ {
		pt, p := nums[b], b
		for {
			p = (p + k) % len(nums)
			nums[p], pt = pt, nums[p]
			if p == b {
				break
			}
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
