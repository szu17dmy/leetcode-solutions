package summary_ranges

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) (r []string) {
	var b, e int
	for b != len(nums) {
		if e < len(nums)-1 && nums[e]+1 == nums[e+1] {
			e++
			continue
		}
		if b == e {
			r = append(r, strconv.Itoa(nums[b]))
		} else {
			r = append(r, fmt.Sprintf("%d->%d", nums[b], nums[e]))
		}
		e++
		b = e
	}
	return
}
