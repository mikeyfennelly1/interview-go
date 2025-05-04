package src

/*
 * Complete the 'minimalHeaviestSetA' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func TwoSum(nums []int, target int) []int {
	ht := make(map[int]int)

	for i, num := range nums {
		// if the compliment for num (target-num) exists in the hash table...
		if j, exists := ht[target-num]; exists {
			// ... return the index of the compliment and num
			return []int{j, i}
		}
		// create an entry in the hashtable for num and its index
		ht[num] = i
	}

	return []int{}
}
