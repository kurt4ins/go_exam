package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}

	first := string(word[0])
	rest := word[1:]
	restRunes := []rune(rest)
	for i, j := 0, len(restRunes)-1; i < j; i, j = i+1, j-1 {
		restRunes[i], restRunes[j] = restRunes[j], restRunes[i]
	}

	return first + string(restRunes)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)
	encryptedWords := make([]string, len(words))

	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	testPhrase1 := "Pepe Schnele is a legend"
	fmt.Printf("Исходная фраза: %s\n", testPhrase1)
	fmt.Printf("Зашифрованная:  %s\n\n", encryptPhrase(testPhrase1))

	testPhrase2 := "Hello World"
	fmt.Printf("Исходная фраза: %s\n", testPhrase2)
	fmt.Printf("Зашифрованная:  %s\n\n", encryptPhrase(testPhrase2))

	testPhrase3 := "Gogo goga"
	fmt.Printf("Исходная фраза: %s\n", testPhrase3)
	fmt.Printf("Зашифрованная:  %s\n\n", encryptPhrase(testPhrase3))

	testPhrase4 := "x y z"
	fmt.Printf("Исходная фраза: %s\n", testPhrase4)
	fmt.Printf("Зашифрованная:  %s\n\n", encryptPhrase(testPhrase4))

	testPhrase5 := "Ganvest is also a legend"
	fmt.Printf("Исходная фраза: %s\n", testPhrase5)
	fmt.Printf("Зашифрованная:  %s\n\n", encryptPhrase(testPhrase5))
}
