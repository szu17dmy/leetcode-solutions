package count_the_number_of_vowel_strings_in_range

func vowelStrings(words []string, left int, right int) int {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	sum := 0
	for i := left; i <= right; i++ {
		word := words[i]
		if vowels[word[:1]] && vowels[word[len(word)-1:]] {
			sum++
		}
	}
	return sum
}
