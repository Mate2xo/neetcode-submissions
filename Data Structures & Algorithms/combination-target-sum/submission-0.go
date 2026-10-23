import (
	"slices"
)

func combinationSum(nums []int, target int) [][]int {
	combinations := [][]int{}

	if len(nums) == 0 {
		return nil
	}

	for i, num := range nums {
		currentCombination := []int{num}
		if num == target {
			combinations = append(combinations, currentCombination)
			return combinations
		}
		if num >= target {
			break
		}
		innerCombinations := combinationSum(nums[i:], target-num)
		for _, combination := range innerCombinations {
			combinations = append(combinations, slices.Concat(currentCombination, combination))
		}
	}
	return combinations
}
