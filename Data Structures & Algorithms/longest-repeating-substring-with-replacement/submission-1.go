func characterReplacement(s string, k int) int {
	longest, charCount := 0, make(map[rune]int, 26)
	for _, char := range s {
		charCount[char]++
		if charCount[char] > longest {
			longest = charCount[char]
		}
	}

	return min(longest + k, len(s))
}
