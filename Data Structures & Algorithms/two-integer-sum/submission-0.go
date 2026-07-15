func twoSum(nums []int, target int) []int {
    num2Index := make(map[int]int)

    for i, v := range nums {
        if idx, ok := num2Index[target-v]; ok {
            return []int{idx, i}
        }

        num2Index[v] = i
    }

    return nil
}