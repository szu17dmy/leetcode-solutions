package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var r [][]int
	r = append(r, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lhs, rhs := r[len(r)-1], intervals[i]
		if lhs[1] < rhs[0] {
			r = append(r, intervals[i])
			continue
		}
		if lhs[1] < rhs[1] {
			r[len(r)-1][1] = rhs[1]
			continue
		}
	}
	return r
}
