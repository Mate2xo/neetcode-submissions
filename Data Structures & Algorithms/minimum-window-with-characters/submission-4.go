func minWindow(s, t string) string {
	if len(s) < len(t) { return "" }

	left, shortest, required  := 0, "", map[byte]int{}

	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}

	for right := 0; right < len(s); right++ {
		letter := s[right]
		if _, found := required[letter]; !found {
			continue
		}

		required[letter]--
		for allCharsPresent(required) {
			if shortest == "" || right-left < len(shortest) {
				if right == len(s) {
					shortest = s[left:len(s)]
				} else {
					shortest = s[left:right+1]
				}
			}
			if _, found := required[s[left]]; found {
				required[s[left]]++
			}
			left++
		}
	}
	return shortest
}

func allCharsPresent(required map[byte]int) bool {
	for _, count := range required {
		if count > 0 {
			return false
		}
	}
	return true
}