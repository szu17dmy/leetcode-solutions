package find_the_longest_balanced_substring_of_a_binary_string

func findTheLongestBalancedSubstring(s string) (r int) {
	cnt := make([]int, 2)
	for i, el := range s {
		if el == '1' {
			cnt[1]++
			r = max(r, 2*min(cnt[0], cnt[1]))
		} else if el == '0' && (i == 0 || s[i-1] == '1') {
			cnt[0] = 1
			cnt[1] = 0
		} else {
			cnt[0]++
		}
	}
	return r
}

func min(l int, r int) int {
	if l > r {
		return r
	}
	return l
}

func max(l int, r int) int {
	if l > r {
		return l
	}
	return r
}
