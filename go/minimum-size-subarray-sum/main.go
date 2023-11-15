package minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	var b, e, sum, min int
	for e != len(nums) {
		sum += nums[e]
		for sum >= target {
			if min == 0 || e-b+1 < min {
				min = e - b + 1
			}
			sum -= nums[b]
			b++
		}
		e++
	}
	return min
}
