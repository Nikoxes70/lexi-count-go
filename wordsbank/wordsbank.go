package wordsbank

import (
	"bufio"
	"fmt"
	"net/http"
)

type repoer interface {
	GetWords() map[string]struct{}
	SetWords(words map[string]struct{})
}

type validatorer interface {
	IsValidWord(word string) bool
}

type WordsBank struct {
	repo repoer
	url  string
	validatorer
}

func NewWordsBank(r repoer, url string, v validatorer) *WordsBank {
	wb := &WordsBank{
		repo:        r,
		url:         url,
		validatorer: v,
	}

	wb.loadWordBank()
	return wb
}

func (wb *WordsBank) IsExists(word string) bool {
	// Rule 3: Word must be part of the word bank.d
	words := wb.repo.GetWords()
	_, exists := words[word]
	return exists
}

func (wb *WordsBank) loadWordBank() {
	resp, err := http.Get(wb.url)
	if err != nil {
		fmt.Printf("WordsBank.loadWordBank http.Get failed - %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("WordsBank.loadWordBank failed - %v\n", err)
	}

	scanner := bufio.NewScanner(resp.Body)
	wordBank := make(map[string]struct{})

	for scanner.Scan() {
		word := scanner.Text()
		if isValid := wb.IsValidWord(word); isValid {
			wordBank[word] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("WordsBank.loadWordBank scanner.Err failed - %v\n", err)
	}

	wb.repo.SetWords(wordBank)
}
