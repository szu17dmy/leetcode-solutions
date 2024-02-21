package plus_one

func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; c != 0 && i >= 0; i-- {
		s := digits[i] + c
		c = s / 10
		digits[i] = s % 10
	}
	if c != 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
