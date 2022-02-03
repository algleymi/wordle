package dictionary

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DictionaryListsWords(t *testing.T) {
	wordlist := GetWordList()
	assert.Equal(t, 12972, len(wordlist))
}

func Test_SelectRandomWordReturnsErrorWhenListIsEmpty(t *testing.T) {
	_, err := SelectRandomWordFromList([]string{})
	assert.True(t, errors.Is(err, ErrEmptyWordlist))
}

func Test_SelectRandomWordReturnsRandomWord(t *testing.T) {
	wordlist := []string{
		"hello",
		"stuff",
		"weary",
		"aahed",
	}

	var uniqueWords []string
	for len(uniqueWords) < len(wordlist) {
		word, _ := SelectRandomWordFromList(wordlist)
		if !contains(uniqueWords, word) {
			uniqueWords = append(uniqueWords, word)
		}
	}

	assert.ElementsMatch(t, uniqueWords, wordlist)
}

func contains(list []string, word string) bool {
	for _, value := range list {
		if word == value {
			return true
		}
	}

	return false
}