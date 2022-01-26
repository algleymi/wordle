package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DictionaryListsWords(t *testing.T) {
	wordlist := GetWordList()
	assert.Len(t, wordlist, 370103)
}

func Test_FiveLetterWordsList(t *testing.T) {
	wordlist := GetFiveLetterWordList()
	assert.Len(t, wordlist, 7186)
}

func Benchmark_FiveLetterWordsList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFiveLetterWordList()
	}
}
