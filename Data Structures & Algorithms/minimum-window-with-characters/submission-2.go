func minWindow(s, t string) string {
	if len(s) < len(t) { return "" }

	left, shortestWindow, required  := 0, [2]int{0, len(s)}, map[byte]int{}

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
			if right-left < shortestWindow[1]-shortestWindow[0] {
				shortestWindow[0], shortestWindow[1] = left, right
			}
			if _, found := required[s[left]]; found {
				required[s[left]]++
			}
			left++
		}
	}
	if shortestWindow[1] == len(s) {
		return s[shortestWindow[0]:]
	} else {
		return s[shortestWindow[0]:shortestWindow[1]+1]
	}
}

func allCharsPresent(required map[byte]int) bool {
	for _, count := range required {
		if count > 0 {
			return false
		}
	}
	return true
}