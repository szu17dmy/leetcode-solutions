package jump_game

func canJump(nums []int) bool {
	d := nums[0]
	if d == 0 && len(nums) > 1 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if i+nums[i] > d {
			d = i + nums[i]
		}
		if d == i && i != len(nums)-1 {
			return false
		}
	}
	return true
}
