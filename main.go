package main

import (
	"cypher/cypher"
	"fmt"
)

func main() {
	fmt.Println("Hello! Welcome to our Cypher tool 👀")
	toEncrypt, encoding, message := cypher.GetInput()
	if toEncrypt == true && encoding == "rot" {
		fmt.Println("Encrypted message using ROT13:", cypher.Rot13(message))
	} else if toEncrypt == false && encoding == "rot" {
		fmt.Println("Decrypted message using ROT13:", cypher.Rot13(message))
	} else if toEncrypt == true && encoding == "custom" {
		fmt.Println("Encrypted message using Simple Shift:", cypher.SimpleShiftEncrypt(message))
	} else if toEncrypt == false && encoding == "custom" {
		fmt.Println("Decrypted message using Simple Shift:", cypher.SimpleShiftDecrypt(message))
	} // else if toEncrypt == true && encoding == "reverse" {
	// 	fmt.Println("Encrypted message using Reverse:", cypher.SimpleShiftDecrypt(message))
	// } else if toEncrypt == false && encoding == "reverse" {
	// 	fmt.Println("Decrypted message using Reverse:", cypher.SimpleShiftDecrypt(message))
	// }
	fmt.Println("Thanks for using our Cypher tool 🥹")
}
