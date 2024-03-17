package _4sum

import "sort"

func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}
			c, d := b+1, len(nums)-1
			for c < d {
				delta := target - nums[d] - nums[c] - nums[b] - nums[a]
				if delta < 0 {
					d--
				} else if delta > 0 {
					c++
				} else {
					if len(ans) > 0 {
						last := ans[len(ans)-1]
						if last[0] == nums[a] && last[1] == nums[b] && last[2] == nums[c] && last[3] == nums[d] {
							c++
							d--
							continue
						}
					}
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
				}
			}
		}
	}
	return ans
}
