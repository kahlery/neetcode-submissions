func hasDuplicate(nums []int) bool {
	n := len(nums)
    seen := make(map[int]struct{}, n)

	for _, v := range nums {		
		if _, exists := seen[v]; exists {
			return true
		}

		seen[v] = struct{}{}
	}
	
	return false
}
