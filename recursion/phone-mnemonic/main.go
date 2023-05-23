package main

import "fmt"

func PhoneNumberMnemonics(phoneNumber string) []string {
	// Write your code here.
	currentMnemonic := make([]byte, len(phoneNumber))
	for i := range currentMnemonic {
		currentMnemonic[i] = '0'
	}
	var mnemonicFound []string

	phoneNumberMnemonicsHelper(0, phoneNumber, currentMnemonic, &mnemonicFound)
	return mnemonicFound
}

var digitLetters = map[byte][]byte{
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
	'0': {'0'},
}

func phoneNumberMnemonicsHelper(index int, phoneNumber string, currentMnemonic []byte, mnemonicFound *[]string) {
	if index == len(phoneNumber) {
		mnemonic := string(currentMnemonic)
		*mnemonicFound = append(*mnemonicFound, mnemonic)
	} else {
		digit := phoneNumber[index]
		letters := digitLetters[digit]
		for _, letter := range letters {
			currentMnemonic[index] = letter
			phoneNumberMnemonicsHelper(index+1, phoneNumber, currentMnemonic, mnemonicFound)
		}
	}
}

func main() {
	fmt.Println(PhoneNumberMnemonics("1905"))
}
