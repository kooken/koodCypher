package main

import (
	"cypher/cypher"
	"fmt"
)

func main() {
	fmt.Println("Hello! Welcome to our Cypher tool 👀")
	toEncrypt, encoding, message := cypher.GetInput()

	if toEncrypt && encoding == "rot" {
		fmt.Println("Encrypted message using ROT13:", cypher.Rot13(message))
	} else if !toEncrypt && encoding == "rot" {
		fmt.Println("Decrypted message using ROT13:", cypher.Rot13(message))
	} else if toEncrypt && encoding == "custom" {
		fmt.Println("Encrypted message using Simple Shift:", cypher.SimpleShiftEncrypt(message))
	} else if !toEncrypt && encoding == "custom" {
		fmt.Println("Decrypted message using Simple Shift:", cypher.SimpleShiftDecrypt(message))
	} else if toEncrypt && encoding == "reverse" {
		fmt.Println("Encrypted message using Reverse:", cypher.ReverseAlphabet(message))
	} else if !toEncrypt && encoding == "reverse" {
		fmt.Println("Decrypted message using Reverse:", cypher.ReverseAlphabet(message))
	}

	fmt.Println("Thanks for using our Cypher tool 🥹")
}
