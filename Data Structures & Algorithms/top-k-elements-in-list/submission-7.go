func topKFrequent(nums []int, k int) []int {
	size := len(nums)
	frequencies := make(map[int]int, size)
	for _, num := range nums {
		frequencies[num]++
	}

	numsByFrequencyIdx := make([][]int, size + 1)
	for num, count := range frequencies {
		numsByFrequencyIdx[count] = append(numsByFrequencyIdx[count], num)
	}

	result := []int{}
	for count := size; count > 0; count-- {
		for _, num := range numsByFrequencyIdx[count] {
			result = append(result, num)
			if len(result) >= k {
				return result
			}
		}
	}
	return result

	// keysByFrequency := slices.SortedFunc(maps.Keys(frequencies), func(a, b int) int {
	// 	return cmp.Compare(b, a)
	// })
	//
	// return keysByFrequency[:k]
	// frequencies := make(map[int]int, len(nums))
	// for _, num := range nums {
	// 	frequencies[num]++
	// }

	// keysByFrequency := slices.SortedFunc(maps.Keys(frequencies), func(a, b int) int {
	// 	return cmp.Compare(frequencies[b], frequencies[a])
	// })

	// return keysByFrequency[:k]
}
