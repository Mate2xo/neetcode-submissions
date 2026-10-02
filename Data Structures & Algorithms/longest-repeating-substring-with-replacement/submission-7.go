func characterReplacement(s string, k int) int {
	charSet, longest := map[byte]bool{}, 0
	for i := 0; i < len(s); i++ {
		charSet[s[i]] = true
	}

	for char := range charSet {
		left, matchedCount := 0, 0

		for right := 0; right < len(s); right++ {
			if s[right] == char {
				matchedCount++
			}
			length := right - left + 1
			if length-matchedCount > k {
				if s[left] == char {
					matchedCount--
				}
				left++
			}
			longest = max(longest, right-left+1)
		}
	}

	return longest
}
