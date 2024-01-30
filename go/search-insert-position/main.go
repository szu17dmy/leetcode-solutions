package search_insert_position

func searchInsert(nums []int, target int) int {
	begin, end := 0, len(nums)
	for begin != end {
		h := (begin + end) >> 1
		if nums[h] > target {
			end = h - 1
		} else if nums[h] < target {
			begin = h + 1
		} else {
			return h
		}
	}
	return end
}
