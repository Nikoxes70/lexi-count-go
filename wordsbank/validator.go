package wordsbank

import (
	"unicode"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) IsValidWord(word string) bool {
	// Rule 1: Word must contain at least 3 characters.
	if len(word) < 3 {
		return false
	}

	// Rule 2: Word must contain only alphabetic characters.
	for _, c := range word {
		if !unicode.IsLetter(c) {
			return false
		}
	}

	return true
}
