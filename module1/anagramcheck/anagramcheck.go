package anagramcheck

// IsAnagram reports whether two strings contain the same characters with the
// same frequencies.
func IsAnagram(first, second string) bool {
	firstRunes := []rune(first)
	secondRunes := []rune(second)

	if len(firstRunes) != len(secondRunes) {
		return false
	}

	characters := make(map[rune]int)
	for _, character := range firstRunes {
		characters[character]++
	}

	for _, character := range secondRunes {
		frequency, present := characters[character]
		if !present {
			return false
		}

		if frequency == 1 {
			delete(characters, character)
		} else {
			characters[character] = frequency - 1
		}
	}

	return len(characters) == 0
}
