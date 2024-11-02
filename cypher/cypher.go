package cypher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() (toEncrypt bool, encoding string, message string) {
	reader := bufio.NewReader(os.Stdin) // Create a reader for standard input.

	// Loop for selecting the operation. This will continue until the user provides valid input.
	for {
		fmt.Printf("Select operation (1/2):\n1. Encrypt.\n2. Decrypt.\n")
		operationInput, _ := reader.ReadString('\n')       // Read input until the user presses Enter
		operationInput = strings.TrimSpace(operationInput) // Trimming spaces from the input

		if operationInput == "1" {
			toEncrypt = true
			fmt.Println("Encryption selected.")
			break
		} else if operationInput == "2" {
			toEncrypt = false
			fmt.Println("Decryption selected.")
			break
		} else {
			fmt.Println("Invalid input! Please enter 1 or 2.")
		}
	}

	// Loop for selecting the cypher. This will continue until the user provides valid input.
	for {
		fmt.Printf("Select cypher (1/2/3):\n1. ROT13.\n2. Reverse.\n3. Simple Shift.\n")
		cypherInput, _ := reader.ReadString('\n')    // Read input until the user presses Enter
		cypherInput = strings.TrimSpace(cypherInput) // Trimming spaces from the input

		if cypherInput == "1" {
			encoding = "rot"
			fmt.Println("ROT13 selected.")
			break
		} else if cypherInput == "2" {
			encoding = "reverse"
			fmt.Println("Reverse selected.")
			break
		} else if cypherInput == "3" {
			encoding = "custom"
			fmt.Println("Simple Shift selected.")
			break
		} else {
			fmt.Println("Invalid input! Please enter 1, 2 or 3.")
		}
	}

	for {
		fmt.Print("Enter the message: ")
		message, _ = reader.ReadString('\n') // Read input until the user presses Enter
		message = strings.TrimSpace(message) // Trim spaces from the input

		if message == "" {
			fmt.Println("Message cannot be empty. Please type something 🥺")
		} else {
			break // Exit loop if input is valid (not empty)
		}
	}

	return toEncrypt, encoding, message
}
