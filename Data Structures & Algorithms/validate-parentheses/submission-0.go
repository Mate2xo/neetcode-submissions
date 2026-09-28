func isValid(s string) bool {
   if len(s) < 2 {
		return false
	}

	chars := []byte(s)

	for chars[0] == chars[len(chars)-1] {
		chars = chars[1 : len(chars)-1]
	}

	if len(chars) != 0 {
		return false
	} else {
		return false
	}
 
}
