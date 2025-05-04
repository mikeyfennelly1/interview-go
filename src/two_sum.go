package src

/*
 * Complete the 'minimalHeaviestSetA' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func TwoSum(nums []int, target int) []int {
	for index1, num1 := range nums {
		compliment := target - num1

		for i := 0; i < len(nums); i++ {
			if i == index1 {
				continue
			}

			if nums[i] == compliment {
				return []int{index1, i}
			}
		}
	}

	return nil
}
