package add_binary

func addBinary(a string, b string) string {
	var p int
	var carrier uint8
	result := ""
	for p < len(a) || p < len(b) {
		var s string
		var n1, n2 uint8 = '0', '0'
		if p < len(a) {
			n1 = a[len(a)-1-p]
		}
		if p < len(b) {
			n2 = b[len(b)-1-p]
		}
		s, carrier = sum(n1, n2, carrier)
		result = s + result
		p++
	}
	if carrier != 0 {
		return "1" + result
	}
	return result
}

func sum(a, b, carrier uint8) (string, uint8) {
	s := (a - '0') + (b - '0') + carrier
	return string(s%2 + '0'), s / 2
}
