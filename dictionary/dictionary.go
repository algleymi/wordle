package dictionary

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"strings"
)

func GetWordList() []string {
	return readWordListFromFile()
}

func GetFiveLetterWordList() []string {
	words := readWordListFromFile()
	var filteredWords []string
	for _, word := range words {
		if len(word) == 5 {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}

func readWordListFromFile() []string {
	data, _ := ioutil.ReadFile("./words_alpha.txt")
	text := string(data)
	words := strings.Split(text, "\n")
	return words
}

func SelectRandomWordFromList(wordlist []string)  (string, error) {
	if len(wordlist) == 0 {
		return "", ErrEmptyWordlist
	}
	index := rand.Intn(len(wordlist))
	return wordlist[index], nil
}

var ErrEmptyWordlist = errors.New("wordlist is empty")
