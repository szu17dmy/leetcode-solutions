package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := x
	var s []int
	for n != 0 {
		s = append(s, n%10)
		n /= 10
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
