package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	p1, p2 := 0, len(numbers)-1
	for numbers[p1]+numbers[p2] != target {
		if numbers[p1]+numbers[p2] < target {
			p1++
		} else {
			p2--
		}
	}
	return []int{p1 + 1, p2 + 1}
}
