package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	t := make([][]int, 2)
	t[0] = make([]int, 26)
	t[1] = make([]int, 26)
	count(ransomNote, t[0])
	count(magazine, t[1])
	for i := 0; i < len(t[0]); i++ {
		if t[0][i] > t[1][i] {
			return false
		}
	}
	return true
}

func count(s string, c []int) {
	for _, el := range s {
		c[el-'a'] += 1
	}
}
