package essaymatcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPairFinder_FindTopWords tests the FindTopWords function of PairFinder
func TestPairFinder_FindTopWords(t *testing.T) {
	pf := NewWordCountPair()

	t.Run("typical usage", func(t *testing.T) {
		wordCounts := map[string]int{
			"word1": 10,
			"word2": 5,
			"word3": 8,
		}

		topWords := pf.FindTopWords(wordCounts, 2)
		assert.Len(t, topWords, 2)
		assert.Equal(t, wordCountPair{"word1", 10}, topWords[0])
		assert.Equal(t, wordCountPair{"word3", 8}, topWords[1])
	})

	t.Run("request more words than available", func(t *testing.T) {
		wordCounts := map[string]int{
			"word1": 3,
			"word2": 1,
		}

		topWords := pf.FindTopWords(wordCounts, 5)
		assert.Len(t, topWords, 2) // Should return only the available words
	})

	t.Run("empty word counts", func(t *testing.T) {
		wordCounts := map[string]int{}

		topWords := pf.FindTopWords(wordCounts, 3)
		assert.Empty(t, topWords)
	})
}
