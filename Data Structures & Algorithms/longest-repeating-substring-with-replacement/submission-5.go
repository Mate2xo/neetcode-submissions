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

		fmt.Printf("left: %d, right: %d - ", left, right)
		left++
		if left > len(s)-1 {
			break
		}
		right--
		// right = left
		fmt.Printf("left: %d, right: %d\n", left, right)
		if s[left] == s[right] {
			// fmt.Println("left == right")
			for s[left] == s[right] && left < len(s)-2 {
				// fmt.Printf("- comparing s[left] %d == s[right] %d\n", s[left], s[right])
				left++
				// right = left
			}
		} else {
			// fmt.Println("outside condition")
			// left++
		}
		// left++
		right = left
		// fmt.Printf("+++++left: %d, right: %d\n", left, right)
	}

	return longest
}
