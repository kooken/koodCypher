package cypher

import "strings"

// Simple Shift Encryption
func SimpleShiftEncrypt(message string) string {
	var result strings.Builder

	for _, char := range message {
		if char >= 'A' && char <= 'Z' {
			// Shift uppercase letters
			if char == 'Z' {
				result.WriteByte('A')
			} else {
				result.WriteByte(byte(char + 1)) // Shift the letter forward by 1
			}
		} else if char >= 'a' && char <= 'z' {
			// Shift lowercase letters
			if char == 'z' {
				result.WriteByte('a')
			} else {
				result.WriteByte(byte(char + 1)) // Shift the letter forward by 1
			}
		} else {
			// Leave non-alphabet characters unchanged
			result.WriteByte(byte(char))
		}
	}
	return result.String()
}

// Simple Shift Decryption
func SimpleShiftDecrypt(encryptedMessage string) string {
	var result strings.Builder

	for _, char := range encryptedMessage {
		if char >= 'A' && char <= 'Z' {
			// Shift uppercase letters
			if char == 'A' {
				result.WriteByte('Z')
			} else {
				result.WriteByte(byte(char - 1)) // Shift the letter backward by 1
			}
		} else if char >= 'a' && char <= 'z' {
			// Shift lowercase letters
			if char == 'a' {
				result.WriteByte('z')
			} else {
				result.WriteByte(byte(char - 1)) // Shift the letter backward by 1
			}
		} else {
			// Leave non-alphabet characters unchanged
			result.WriteByte(byte(char))
		}
	}
	return result.String()
}
