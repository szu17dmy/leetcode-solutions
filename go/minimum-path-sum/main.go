package minimum_path_sum

import "math"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	dp[0] = make([]int, len(grid[0])+1)
	for i := 0; i < len(grid[0])+1; i++ {
		dp[0][i] = math.MaxInt
	}
	dp[1] = make([]int, len(grid[0])+1)
	dp[1][0] = math.MaxInt
	dp[1][1] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[1][i+1] = dp[1][i] + grid[0][i]
	}
	for i := 2; i <= len(grid); i++ {
		dp[i] = make([]int, len(grid[i-1])+1)
		dp[i][0] = 20000
		dp[i][1] = dp[i-1][1] + grid[i-1][0]
		for j := 2; j <= len(grid[i-1]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[len(grid)][len(grid[len(grid)-1])]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
