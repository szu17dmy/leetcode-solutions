package maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	return getMaxBlockFromCuts(horizontalCuts, h) * getMaxBlockFromCuts(verticalCuts, w) % 1_000_000_007
}

func getMaxBlockFromCuts(cuts []int, length int) int {
	sort.Ints(cuts)
	r := cuts[0]
	for i := 1; i < len(cuts); i++ {
		r = max(r, cuts[i]-cuts[i-1])
	}
	return max(r, length-cuts[len(cuts)-1])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
