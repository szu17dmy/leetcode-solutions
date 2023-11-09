package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[cur-1] == nums[i] {
			continue
		}
		nums[cur] = nums[i]
		cur++
	}
	return cur
}
