package longest_palindromic_substring

func longestPalindrome(s string) (ans string) {
	var max int
	for i := 0; i < len(s); i++ {
		for j := 0; 0 <= i-j && i+j < len(s); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			if 2*j+1 > max {
				max, ans = 2*j+1, s[i-j:i+j+1]
			}
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			for j := 0; 0 <= i-j && i+1+j < len(s); j++ {
				if s[i-j] != s[i+1+j] {
					break
				}
				if 2*j+2 > max {
					max, ans = 2*j+2, s[i-j:i+1+j+1]
				}
			}
		}
	}
	return ans
}
