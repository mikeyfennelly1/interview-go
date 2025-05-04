package src

func TwoSum2(nums []int, target int) []int {
	ht := make(map[int]int)

	for i, num := range nums {
		if j, exists := ht[target-num]; exists {
			return []int{j, i}
		}

		ht[num] = i
	}

	return nil
}
