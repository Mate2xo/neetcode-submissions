func lengthOfLongestSubstring(s string) int {
	longest := 0
	if len(s) == 0 {
		return longest
	}

	left, lastCharIdx := 0, map[byte]int{}
	for right :=0; right<len(s);right++{
		fmt.Printf("left %d(%c), right %d(%c), map %v, longest %d(%s)\n", left, s[left], right, s[right], lastCharIdx, longest, s[left:right])
		letter:= s[right]
		if _, found := lastCharIdx[letter]; found {
			left++
			if s[left] == s[right] && left < right { left++}
		} 
		currentLength := right - left + 1
		longest = int(math.Max(float64(currentLength), float64(longest)))
		fmt.Printf("currentLength: %d\n", currentLenght)
		lastCharIdx[letter] = right
	}

	return longest
}
