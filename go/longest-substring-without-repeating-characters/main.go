package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) (ans int) {
	set := map[byte]bool{}
	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := set[s[r]]; !ok {
			set[s[r]] = true
			if r-l+1 > ans {
				ans = r - l + 1
			}
		} else {
			for {
				if s[l] == s[r] {
					l++
					break
				}
				delete(set, s[l])
				l++
			}
		}
	}
	return ans
}
