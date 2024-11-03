package cypher

import "strings"

// function to reverse alphabetic characters using rune
func ReverseAlphabet(s string) string {
	var result strings.Builder
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			// Reverse lowercase letters: 'a' <-> 'z', 'b' <-> 'y', etc.
			newChar := 'z' - (char - 'a')
			result.WriteRune(newChar)
		} else if char >= 'A' && char <= 'Z' {
			// Reverse uppercase letters: 'A' <-> 'Z', 'B' <-> 'Y', etc.
			newChar := 'Z' - (char - 'A')
			result.WriteRune(newChar)
		} else {
			// Non-alphabet characters are unchanged
			result.WriteRune(char)
		}
	}
	return result.String()
}
