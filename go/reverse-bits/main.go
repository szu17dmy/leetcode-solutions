package reverse_bits

func reverseBits(num uint32) (r uint32) {
	for i := 0; i < 32; i++ {
		r <<= 1
		r |= (num >> i) & 1
	}
	return r
}
