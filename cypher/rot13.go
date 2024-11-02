package cypher

import "strings"

// ROT13 encryption/decryption
func Rot13(s string) string {
	var result strings.Builder // Create a strings.Builder for the result string
	for _, r := range s {      // Iterate over each rune (character) in the input string
		if 'a' <= r && r <= 'z' { // Check if the character is a lowercase letter
			result.WriteRune('a' + (r-'a'+13)%26) // Shift the letter and wrap around if necessary
		} else if 'A' <= r && r <= 'Z' { // Check if the character is an uppercase letter
			result.WriteRune('A' + (r-'A'+13)%26) // Shift the letter and wrap around if necessary
		} else {
			// If the character is not a letter, write it unchanged to the result
			result.WriteRune(r)
		}
	}
	// Return the resulting string after processing all characters
	return result.String()
}
