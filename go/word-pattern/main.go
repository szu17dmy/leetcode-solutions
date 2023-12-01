package word_pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	chars := strings.Split(pattern, "")
	words := strings.Split(s, " ")
	if len(chars) != len(words) {
		return false
	}
	m := map[string]string{}
	mr := map[string]string{}
	for i := 0; i < len(chars); i++ {
		if word, ok := m[chars[i]]; ok {
			if word != words[i] {
				return false
			}
		} else if ch, ok := mr[words[i]]; ok {
			if ch != chars[i] {
				return false
			}
		} else {
			m[chars[i]] = words[i]
			mr[words[i]] = chars[i]
		}
	}
	return true
}
