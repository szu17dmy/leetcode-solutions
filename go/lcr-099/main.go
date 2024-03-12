package lcr_099

import "math"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0])+1)
		dp[i][0] = math.MaxInt
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = math.MaxInt
	}
	dp[1][1] = grid[0][0]
	for i := 2; i < len(dp[1]); i++ {
		dp[1][i] = dp[1][i-1] + grid[0][i-1]
	}
	for i := 2; i < len(dp); i++ {
		dp[i][1] = dp[i-1][1] + grid[i-1][0]
	}
	for i := 2; i < len(dp); i++ {
		for j := 2; j < len(dp[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	lastRow := dp[len(dp)-1]
	return lastRow[len(lastRow)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
