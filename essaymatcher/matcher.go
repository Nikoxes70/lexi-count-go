package essaymatcher

import (
	"regexp"
	"strings"
)

type validatorer interface {
	IsValidWord(word string) bool
}

type scrapperer interface {
	Scrap(url string, attempt int) (string, error)
}

type wordsBank interface {
	IsExists(word string) bool
}

type EssayMatcher struct {
	validator validatorer
	scrapper  scrapperer
	wordsBank
}

func NewEssayMatcher(v validatorer, s scrapperer, wb wordsBank) *EssayMatcher {
	return &EssayMatcher{
		v,
		s,
		wb,
	}
}

func (em *EssayMatcher) ProcessEssay(url string, wordCounts map[string]int) error {
	txt, err := em.scrapper.Scrap(url, 1)
	if err != nil {
		return err
	}

	newWords := extractWords(txt)
	for _, word := range newWords {
		if isValid := em.validator.IsValidWord(word); isValid {
			if isExist := em.wordsBank.IsExists(word); isExist {
				wordCounts[word]++
			}
		}
	}

	return nil
}

func extractWords(content string) []string {
	re := regexp.MustCompile(`\b[a-zA-Z]{3,}\b`)
	return re.FindAllString(strings.ToLower(content), -1)
}
