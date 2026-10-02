func characterReplacement(s string, k int) int {
	left, longest, charFreq, maxFreq := 0, 0, map[byte]int{}, 0

	for right := 0; right < len(s); right++ {
		letter := s[right]
		charFreq[letter]++
		if charFreq[letter] > maxFreq {
			maxFreq = charFreq[letter]
		}

		windowSize := right - left +1
		if windowSize-maxFreq > k {
			charFreq[s[left]]--
			left++
		}

		longest = max(longest, right-left+1)
	}
	return longest
}
