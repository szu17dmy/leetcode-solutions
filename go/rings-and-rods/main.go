package rings_and_rods

func countPoints(rings string) int {
	rbs := make([]byte, 10)
	m := map[byte]byte{
		'R': 0b1,
		'G': 0b10,
		'B': 0b100,
	}
	for i := 0; i < len(rings)/2; i++ {
		rbs[rings[2*i+1]-'0'] |= m[rings[2*i]]
	}
	sum := 0
	for i := 0; i < len(rbs); i++ {
		if rbs[i] == 0b111 {
			sum++
		}
	}
	return sum
}
