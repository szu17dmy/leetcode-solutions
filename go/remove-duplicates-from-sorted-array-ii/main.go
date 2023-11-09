package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	var cur int
	for i := 0; i < len(nums); i++ {
		if i > 1 && nums[cur-2] == nums[i] && nums[cur-1] == nums[i] {
			continue
		}
		nums[cur] = nums[i]
		cur++
	}
	return cur
}
