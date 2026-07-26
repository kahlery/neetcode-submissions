func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	prefix := make([]int, n)
	suffix := make([]int, n)

	for i, _ := range nums {
		if i == 0 {
			prefix[0] = nums[0]
			continue
		}
		prefix[i] = prefix[i-1] * nums[i]
	}

	fmt.Println("prefix: ", prefix)

	for i := n-1; i >= 0; i-- {
		if i == n-1 {
			suffix[n-1] = nums[n-1]
			continue
		}
		suffix[i] = suffix[i+1] * nums[i]
	}

	fmt.Println("suffix: ", suffix)

	for i, _ := range nums {
		if i == 0 {
			res[i] = suffix[i+1]
		} else if i == n-1 {
			res[i] = prefix[i-1]
 		} else {
			res[i] = prefix[i-1] * suffix[i+1]
		}
	}

	fmt.Println("res: ", res)

	return res
}
