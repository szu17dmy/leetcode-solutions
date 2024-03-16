package container_with_most_water

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	area := min(height[left], height[right]) * (right - left)
	ans = area
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		area = min(height[left], height[right]) * (right - left)
		if area > ans {
			ans = area
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
