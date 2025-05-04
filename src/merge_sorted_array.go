package src

func Merge(nums1 *[]int, m *int, nums2 *[]int, n *int) {
	// if the amount of numbers left to be processed in nums2 (i.e. n)
	// is 0, then we can terminate
	if *n == 0 {
		return
	}

	cur := (*nums2)[*n-1]
	for i := *m; i >= 0; i-- {
		// move the value up one, as we expect to find a condition match
		secondLastIndex := len(*nums1) - 2
		for j := secondLastIndex; j >= 0; j-- {
			// the next value (ie the one at index + 1) equals
			// the current nums2 item.
			if j < len(*nums1) {
				(*nums1)[j+1] = (*nums1)[j]
			}
		}
		if (*nums1)[i] < cur {
			// add the current nums2 item (cur) to the nums1 array
			(*nums1)[i] = (*nums2)[*n-1]
		}
		// in any other case, stay moving down nums1
	}

	// decrement the value of n
	*n--

	// run the algorithm again
	Merge(nums1, m, nums2, n)
}
