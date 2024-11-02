# Cypher Tool 🔐

This command-line tool allows users to **encrypt** and **decrypt** messages using various encryption techniques. The tool offers a simple interface to choose an operation, select an encryption type, and input a message. Non-alphabet characters are preserved during encryption and decryption.

## Features

- **Encrypt** or **Decrypt** any message.
- Supports three encryption types:
  - **ROT13**: Shifts each letter by 13 positions in the alphabet.
  - **Reverse Alphabet**: Replaces each letter with its "mirrored" letter (A ↔ Z, B ↔ Y, etc.).
  - **Simple Shift**: Each letter is shifted one position forward for encryption and one position backward for decryption.
- **User-friendly input validation**: If an input is invalid, the tool prompts the user to re-enter until a valid input is provided.

## Usage

1. **Run the program**:
   go run main.go

2. **Follow the prompts**:
   - **Select the Operation**: Enter `1` for Encrypt or `2` for Decrypt.
   - **Choose the Cypher Type**:
     - **1** for ROT13
     - **2** for Reverse Alphabet
     - **3** for the Simple Shift
   - **Enter the Message** you want to encrypt or decrypt.

3. **Receive the Output**: The tool will display the encrypted or decrypted message.

### Example

Hello! Welcome to our Cypher tool 👀

Select operation (1/2):
1. Encrypt.
2. Decrypt.
> 1

Select cypher (1/2/3):
1. ROT13.
2. Reverse Alphabet.
3. Simple Shift.
> 1

Enter the message:
> Hello, World!

Encrypted message using ROT13: Uryyb, Jbeyq!
Thanks for using our Cypher tool 🥹

## Cyphers Used

1. **ROT13**: Each letter is shifted by 13 places in the alphabet, wrapping around if necessary. For example, `A` becomes `N`, and `N` becomes `A`. This cypher is symmetrical, meaning encrypting and decrypting are the same operation.

2. **Reverse Alphabet**: Each letter is "mirrored" across the alphabet center (A ↔ Z, B ↔ Y, C ↔ X, etc.). For example, `A` becomes `Z`, and `B` becomes `Y`.

3. **Simple Shift**: This method shifts each letter in the message one position forward for encryption and one position backward for decryption. For example, `A` becomes `B`, `B` becomes `C`, ..., `Z` wraps around to become `A`. The same applies for lowercase letters, where `a` becomes `b`, and so on. Non-alphabetic characters remain unchanged.


---

Enjoy the Cypher Tool! 🎉

By Mariia, Muhammed & Abbas
