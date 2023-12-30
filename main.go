package main

import (
	"fmt"
	"strings"
)

type rotor struct {
	wiring   string
	position int
}

type EnigmaMachine struct {
	plugBoard map[rune]rune
}

var plugboardKeys = map[rune]rune{
	'A': 'E', 'B': 'K', 'C': 'M', 'D': 'F', 'E': 'L', 'F': 'G',
	'G': 'D', 'H': 'Q', 'I': 'V', 'J': 'Z', 'K': 'N', 'L': 'T',
	'M': 'O', 'N': 'W', 'O': 'Y', 'P': 'H', 'Q': 'X', 'R': 'U',
	'S': 'S', 'T': 'P', 'U': 'A', 'V': 'I', 'W': 'B', 'X': 'R',
	'Y': 'C', 'Z': 'J',
}

var reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var plugboardkeys = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"

func main() {

	plugboard := NewEnigmaMachine(plugboardKeys)

	rotorI := rotor{wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ"}
	rotorII := rotor{wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE"}
	rotorIII := rotor{wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO"}
	rotorI.position = 0
	rotorII.position = 0
	rotorIII.position = 0

	plaintext := "HELLO"
	encryptedText := plugboard.enigmaEncrypt(plaintext, rotorI, rotorII, rotorIII)
	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Encrypted Text: ", encryptedText)

	rotorI.position = 0
	rotorII.position = 0
	rotorIII.position = 0
	decryptedText := plugboard.enigmaDecrypt(encryptedText, rotorI, rotorII, rotorIII)
	fmt.Println("Decrypted Text: ", decryptedText)
}

func (e *EnigmaMachine) enigmaEncrypt(plaintext string, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {

			// Rotate rotors before encryption
			rotateRotors(&rotors)

			// Pass the character through the plugboard
			char = plugboardSubstitution(char, e.plugBoard)

			// Pass the character through the rotors from right to left
			char = substitute(char, rotors[2])
			char = substitute(char, rotors[1])
			char = substitute(char, rotors[0])

			// Pass the character through the reflector
			char = reflector(char)

			// Pass the character through the rotors from left to right
			char = substitute(char, rotors[0])
			char = substitute(char, rotors[1])
			char = substitute(char, rotors[2])

			// Pass the character through the plugboard
			char = plugboardSubstitution(char, e.plugBoard)

			encrypted.WriteRune(char)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func (e *EnigmaMachine) enigmaDecrypt(plaintext string, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var decrypted strings.Builder

	reverseKey := make(map[rune]rune)
	for k, v := range plugboardKeys {
		reverseKey[v] = k

	}

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Rotate rotors before encryption
			rotateRotors(&rotors)

			// Pass the character through the plugboard
			char = plugb(char)

			char = decrypt(char, rotors[2])
			char = decrypt(char, rotors[1])
			char = decrypt(char, rotors[0])
			char = reflector(char)
			char = decrypt(char, rotors[0])
			char = decrypt(char, rotors[1])
			char = decrypt(char, rotors[2])

			// Pass the character through the plugboard
			char = plugb(char)

			decrypted.WriteRune(char)
		} else {
			// Non-alphabetic characters are not modified
			decrypted.WriteRune(char)
		}
	}

	return decrypted.String()
}

func rotateRotors(rotors *[]rotor) {
	for i := 0; i < len(*rotors); i++ {
		(*rotors)[i].position++
		if (*rotors)[i].position >= 26 {
			(*rotors)[i].position = 0
		}
	}
}

func substitute(char rune, rotor rotor) rune {
	index := (int(char-'A') + rotor.position) % 26
	return rune(rotor.wiring[index])
}

func decrypt(char rune, rotor rotor) rune {
	index := (strings.IndexRune(rotor.wiring, char) - rotor.position + 26) % 26
	return rune(alphabet[index])
}

func reflector(char rune) rune {
	index := strings.IndexRune(reflectorB, char)
	return rune(alphabet[index])
}

func plugb(char rune) rune {
	index := strings.IndexRune(plugboardkeys, char)
	return rune(alphabet[index])
}

func NewEnigmaMachine(plugboard map[rune]rune) *EnigmaMachine {
	return &EnigmaMachine{
		plugBoard: plugboard,
	}
}

func plugboardSubstitution(char rune, plugboard map[rune]rune) rune {
	if plug, ok := plugboard[char]; ok {
		return plug
	}
	return char
}
