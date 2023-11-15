package valid_palindrome

func isPalindrome(s string) bool {
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		if !isWordChar(s[p1]) {
			p1++
			continue
		}
		if !isWordChar(s[p2]) {
			p2--
			continue
		}
		if !equals(s[p1], s[p2]) {
			return false
		}
		p1++
		p2--
	}
	return true
}

func isWordChar(c byte) bool {
	if ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}
	return false
}

func equals(a, b byte) bool {
	if 'a' <= a && a <= 'z' {
		a -= 'a' - 'A'
	}
	if 'a' <= b && b <= 'z' {
		b -= 'a' - 'A'
	}
	return a == b
}
