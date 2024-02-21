package triangle

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := [][]int{
		triangle[0],
	}
	var m int
	for i := 1; i < len(triangle); i++ {
		dp = append(dp, make([]int, i+1))
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		m = dp[i][0]
		for j := 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			if dp[i][j] < m {
				m = dp[i][j]
			}
		}
		dp[i][len(triangle[i])-1] = dp[i-1][len(triangle[i])-2] + triangle[i][len(triangle[i])-1]
		if dp[i][len(triangle[i])-1] < m {
			m = dp[i][len(triangle[i])-1]
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
