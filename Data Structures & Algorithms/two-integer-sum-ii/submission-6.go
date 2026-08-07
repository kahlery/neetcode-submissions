func twoSum(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left+1, right+1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
