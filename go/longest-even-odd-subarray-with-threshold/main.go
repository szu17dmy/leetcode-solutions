package longest_even_odd_subarray_with_threshold

func longestAlternatingSubarray(nums []int, threshold int) int {
	var l, ans int
	for i := 0; i < len(nums); i++ {
		if nums[i] > threshold {
			l = 0
			continue
		}
		if l != 0 {
			if nums[i-1]%2 != nums[i]%2 {
				l++
				if l > ans {
					ans = l
				}
				continue
			}
			l = 0
		}
		if nums[i]%2 == 0 {
			l = 1
			if l > ans {
				ans = l
			}
		}
	}
	return ans
}
