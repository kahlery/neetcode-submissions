func twoSum(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n-1

	sum := 0
	for left < right {
		sum = nums[left] + nums[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left+1, right+1}
		}
	}

	return nil
}
