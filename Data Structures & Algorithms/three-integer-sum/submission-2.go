import "slices"

func threeSum(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)

	if n < 3 {
		return [][]int{}
	}

	res := [][]int{}

	for left := 0; left < n-2; left++ {
		// Skip duplicate values for the first number.
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		// Since the array is sorted, no solution is possible.
		if nums[left] > 0 {
			break
		}

		mid := left + 1
		right := n - 1

		for mid < right {
			sum := nums[left] + nums[mid] + nums[right]

			if sum < 0 {
				mid++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{
					nums[left],
					nums[mid],
					nums[right],
				})

				// Skip duplicates for mid and right.
				for mid < right && nums[mid] == nums[mid+1] {
					mid++
				}

				for mid < right && nums[right] == nums[right-1] {
					right--
				}

				mid++
				right--
			}
		}
	}

	return res
}