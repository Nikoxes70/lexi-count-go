package essaymatcher

import (
	"sort"
)

type wordCountPair struct {
	Word  string
	Count int
}

type PairFinder struct{}

func NewWordCountPair() *PairFinder {
	return &PairFinder{}
}

func (pf *PairFinder) FindTopWords(wordCounts map[string]int, n int) []wordCountPair {
	pairs := make([]wordCountPair, 0, len(wordCounts))
	for word, count := range wordCounts {
		pairs = append(pairs, wordCountPair{Word: word, Count: count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	if n > len(pairs) {
		n = len(pairs)
	}
	return pairs[:n]
}
