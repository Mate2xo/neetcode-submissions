import "slices"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	openers := []rune{}
	matchingOpener := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if slices.Contains([]rune{'(', '[', '{'}, char) {
			openers = append(openers, char)
		} else {
			opener := openers[len(openers)-1]
			openers = openers[:len(openers)-1]
			if opener != matchingOpener[char] {
				return false
			}
		}
	}

	return true
}
