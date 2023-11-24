package length_of_last_word

func lengthOfLastWord(s string) int {
	var pb, pe int
	for pb = len(s) - 1; pb >= 0 && s[pb] == ' '; pb-- {
	}
	for ; pb >= 0 && isChar(s[pb]); pb-- {
	}
	for pe = len(s) - 1; pe >= 0 && s[pe] == ' '; pe-- {
	}
	return pe - pb
}

func isChar(ch byte) bool {
	if ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9') {
		return true
	}
	return false
}
