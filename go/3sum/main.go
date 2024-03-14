package _3sum

import "sort"

func threeSum(nums []int) (r [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				if !(len(r) != 0 && r[len(r)-1][0] == nums[i] && r[len(r)-1][1] == nums[j]) {
					r = append(r, []int{nums[i], nums[j], nums[k]})
				}
				j++
				k--
			}
		}
	}
	return r
}
