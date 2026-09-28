func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	chars := []byte(s)

	for len(chars) != 0 {
		opening, closing := string(chars[0]), string(chars[len(chars)-1])
		switch opening {
		case "(":
			if closing != ")" {
				return false
			}
		case "[":
			if closing != "]" {
				return false
			}
		case "{":
			if closing != "}" {
				return false
			}
		}
		chars = chars[1 : len(chars)-1]
	}

	return true
}
