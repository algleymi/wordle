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

func readWordListFromFile() []string {
	data, _ := ioutil.ReadFile("./words.txt")
	text := string(data)
	words := strings.Split(text, "\n")
	return words
}

func SelectRandomWordFromList(generator *rand.Rand, wordlist []string) (string, error) {
	if len(wordlist) == 0 {
		return "", ErrEmptyWordlist
	}
	index := generator.Intn(len(wordlist))

	return SelectWordFromList(wordlist, index), nil
}

func SelectWordFromList(wordlist []string, index int) string {
	return wordlist[index]
}

var ErrEmptyWordlist = errors.New("wordlist is empty")
