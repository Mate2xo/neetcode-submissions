func isPalindrome(s string) bool {
	alphanumeric := strings.ToLower(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		} else {
			return -1
		}
	}, s))

	fmt.Printf("string: %s", alphanumeric)
	for start, end := 0, len(alphanumeric)-1; start < end; {
		if alphanumeric[start] != alphanumeric[end] {
			return false
		}
		start++
		end--
	}
	return true
}
