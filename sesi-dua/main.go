package main

import "fmt"

func countWord(word string) map[string]int {

	// change alphabet into rune type so
	var runeLength = map[rune]int{}
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
		char := rune(word[i])
		runeLength[char]++
	}

	// reverse rune type into string
	var wordLength = map[string]int{}
	for k, v := range runeLength {
		newKey := string(k)
		wordLength[newKey] = v
	}

	return wordLength
}

func main() {
	wordLength := countWord("mata angin")
	fmt.Println(wordLength)
}
