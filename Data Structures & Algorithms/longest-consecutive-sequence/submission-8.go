import "slices"

func longestConsecutive(nums []int) int {
	slices.Sort(nums)
	rate, maxRate := 1, 0

 	for i, v := range nums {
		if i == 0 {
			maxRate = 1
			continue
		} else if v == nums[i-1] {
			continue
		} else if v == nums[i-1]+1 {
			rate++
			maxRate = max(rate, maxRate)
		} else {
			rate = 1
		}
	}

	return maxRate
}
