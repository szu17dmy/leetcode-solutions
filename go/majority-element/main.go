package majority_element

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] += 1
	}
	max, r := 0, nums[0]
	for k, v := range m {
		if m[k] > max {
			max = v
			r = k
		}
	}
	return r
}
