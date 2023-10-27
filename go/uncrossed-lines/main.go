package uncrossed_lines

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	mat := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		mat[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				mat[i][j] = mat[i-1][j-1] + 1
			} else {
				mat[i][j] = max(mat[i][j-1], mat[i-1][j])
			}
		}
	}
	return mat[len(nums1)][len(nums2)]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
