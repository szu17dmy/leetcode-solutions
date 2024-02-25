package trapping_rain_water

func trap(height []int) int {
	lmax := make([]int, len(height))
	lmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}
	rmax := make([]int, len(height))
	rmax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}
	var sum int
	for i := 0; i < len(height); i++ {
		sum += min(lmax[i], rmax[i]) - height[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
