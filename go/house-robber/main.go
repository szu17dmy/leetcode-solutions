package house_robber

func rob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	if nums[0] < nums[1] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	if len(nums) == 2 {
		return dp[1]
	}
	for i := 2; i < len(nums); i++ {
		if dp[i-2]+nums[i] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
