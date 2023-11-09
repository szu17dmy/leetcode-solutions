package remove_element

func removeElement(nums []int, val int) int {
	var cur int
	for i, num := range nums {
		if num == val {
			continue
		}
		nums[cur] = nums[i]
		cur++
	}
	return cur
}
