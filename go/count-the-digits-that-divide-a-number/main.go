package count_the_digits_that_divide_a_number

func countDigits(num int) (r int) {
	n := num
	for n != 0 {
		if num%(n%10) == 0 {
			r++
		}
		n /= 10
	}
	return r
}
