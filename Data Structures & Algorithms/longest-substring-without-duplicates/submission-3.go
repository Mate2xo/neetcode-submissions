func lengthOfLongestSubstring(s string) int {
	longest := 0
	if len(s) == 0 {
		return longest
	}

	left, lastCharIdx := 0, make(map[byte]int, len(s))
	for right := 0; right < len(s); right++ {
		letter := s[right]
		if idx, found := lastCharIdx[letter]; found {
			left = max(idx + 1, left)
		}
		currentLength := right - left + 1
		longest = max(currentLength, longest)
		lastCharIdx[letter] = right
	}

	return longest
}
