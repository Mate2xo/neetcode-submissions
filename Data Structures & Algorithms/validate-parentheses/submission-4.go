func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	openers := []rune{}

	for _, char := range s {
		switch string(char) {
		case "(", "[", "{":
			openers = append(openers, char)
		case ")":
			if len(openers) == 0 { return false }
			opener := openers[len(openers)-1]
			openers = openers[:len(openers)-1]
			if string(opener) != "(" {
				return false
			}
		case "]":
			if len(openers) == 0 { return false }
			opener := openers[len(openers)-1]
			openers = openers[:len(openers)-1]
			if string(opener) != "[" {
				return false
			}
		case "}":
			if len(openers) == 0 { return false }
			opener := openers[len(openers)-1]
			openers = openers[:len(openers)-1]
			if string(opener) != "{" {
				return false
			}
		}
	}

	return len(openers) == 0
}
