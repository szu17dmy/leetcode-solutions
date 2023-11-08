package repeated_dna_sequences

const sequenceLength = 10

var dnaBin = map[byte]int{
	'A': 0b00,
	'C': 0b01,
	'G': 0b10,
	'T': 0b11,
}

func findRepeatedDnaSequences(s string) (result []string) {
	if len(s) < 10 {
		return nil
	}
	var f int
	for i := 0; i < sequenceLength; i++ {
		f = (f << 2) | dnaBin[s[i]]
	}
	m := map[int]int{
		f: 1,
	}
	for i := 1; i <= len(s)-sequenceLength; i++ {
		f = (f<<2)&((1<<20)-1) | dnaBin[s[i+9]]
		if m[f] == 1 {
			result = append(result, s[i:i+10])
		}
		m[f]++
	}
	return result
}
