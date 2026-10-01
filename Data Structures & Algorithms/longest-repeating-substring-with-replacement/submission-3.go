func characterReplacement(s string, k int) int {
	left, right, longest := 0, 0, 0

	for left < len(s) {
		remainingReplacements, length := k, 0
		for right < len(s) && (s[left] == s[right] || remainingReplacements > 0) {
			if s[left] != s[right] {
				remainingReplacements--
			}
			right++
			length++
		}
		for remainingReplacements !=0 && length <= len(s) {
			length = min(length + 1, len(s))
			remainingReplacements--
		}
		longest = max(longest, length)

		left++
		right = left
	}

	return longest
}
