package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) || len(s) == 0 {
		return false
	}
	m1 := map[byte]byte{}
	m2 := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if c, ok := m1[s[i]]; ok {
			if c != t[i] {
				return false
			}
		} else {
			m1[s[i]] = t[i]
		}
		if c, ok := m2[t[i]]; ok {
			if c != s[i] {
				return false
			}
		} else {
			m2[t[i]] = s[i]
		}
	}
	return true
}
