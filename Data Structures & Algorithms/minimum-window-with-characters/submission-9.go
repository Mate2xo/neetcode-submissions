type matchingLetters struct {
	required map[byte]int
	have     int
	need     int
	leftP    int
	shortest string
}

func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	state := matchingLetters{map[byte]int{}, 0, 0, 0, ""}
	for i := 0; i < len(t); i++ {
		state.required[t[i]]++
		state.need++
	}

	for rightP := 0; rightP < len(s); rightP++ {
		letter := s[rightP]
		if _, found := state.required[letter]; !found {
			continue
		}

		trackMatchedLetters(&state, letter)
		for allCharsPresent(&state) {
			if state.shortest == "" || rightP-state.leftP < len(state.shortest) {
				updateShortestString(&state, s, rightP)
			}
			incrementLeftPointer(&state, s[state.leftP])
		}
	}

	return state.shortest
}

func trackMatchedLetters(state *matchingLetters, letter byte) {
	state.required[letter]--
	if state.required[letter] <= 0 {
		state.have++
	}
}

func allCharsPresent(state *matchingLetters) bool {
	return state.have >= state.need
}

func updateShortestString(state *matchingLetters, s string, rightP int) {
	if rightP == len(s) {
		state.shortest = s[state.leftP:]
	} else {
		state.shortest = s[state.leftP : rightP+1]
	}
}

func incrementLeftPointer(state *matchingLetters, leftLetter byte) {
	if _, found := state.required[leftLetter]; found {
		state.required[leftLetter]++
		state.have--
	}
	state.leftP++
}
