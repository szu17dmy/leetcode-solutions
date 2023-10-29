package h_index

import "sort"

func hIndex(citations []int) (h int) {
	sort.Sort(sort.IntSlice(citations))
	for i := len(citations); i != 0 && citations[i-1] > h; i-- {
		h++
	}
	return h
}
