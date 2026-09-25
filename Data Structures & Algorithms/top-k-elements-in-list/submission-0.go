import "slices"
import "maps"
import "cmp"

func topKFrequent(nums []int, k int) []int {
	frequencies := make(map[int]int, len(nums))
	for _, num := range nums {
		frequencies[num]++
	}

	keysByFrequency := slices.SortedFunc(maps.Keys(frequencies), func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return keysByFrequency[:k]
}
