package dictionary

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DictionaryListsWords(t *testing.T) {
	wordlist := GetWordList()
	assert.Equal(t, 12949, len(wordlist))
}

func Test_SelectRandomWordReturnsErrorWhenListIsEmpty(t *testing.T) {
	generator := rand.New(rand.NewSource(1))
	_, err := SelectRandomWordFromList(generator, []string{})
	assert.True(t, errors.Is(err, ErrEmptyWordlist))
}

func Test_SelectRandomWordReturnsRandomWord(t *testing.T) {
	wordlist := []string{
		"hello",
		"stuff",
		"weary",
		"aahed",
	}

	generator := rand.New(rand.NewSource(1))
	word, _ := SelectRandomWordFromList(generator, wordlist)

	assert.Equal(t, "stuff", word)
}
