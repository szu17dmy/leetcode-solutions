package maximum_binary_string_after_change

func maximumBinaryString(binary string) string {
	buf := []byte(binary)
	zl := 1
	for i := 0; i < len(buf)-1; i++ {
		if buf[i] == '1' {
			continue
		}
		if zl <= i {
			zl = i+1
		}
		for ; zl < len(buf); zl++ {
			if buf[zl] == '0' {
				buf[i+1], buf[zl] = buf[zl], buf[i+1]
				buf[i] = '1'
				zl++
				break
			}
		}
	}
	return string(buf)
}
