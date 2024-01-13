package wordsbank

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const invalidWord = "inv@l1d"

// MockRepoer mocks the repoer interface
type MockRepoer struct {
	mock.Mock
	words map[string]struct{}
}

func (m *MockRepoer) GetWords() map[string]struct{} {
	return m.words
}

func (m *MockRepoer) SetWords(words map[string]struct{}) {
	m.words = words
}

// MockValidatorer mocks the validatorer interface
type MockValidatorer struct {
	mock.Mock
}

func (m *MockValidatorer) IsValidWord(word string) bool {
	args := m.Called(word)
	if word == invalidWord {
		return false
	}
	return args.Bool(0)
}

// TestWordsBank_IsExists tests the IsExists function of WordsBank
func TestWordsBank_IsExists(t *testing.T) {
	mockRepo := &MockRepoer{words: make(map[string]struct{})}
	mockValidator := new(MockValidatorer)
	mockValidator.On("IsValidWord", mock.Anything).Return(true)

	// Set up a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("word1\nword2\nword3"))
	}))
	defer ts.Close()

	wb := NewWordsBank(mockRepo, ts.URL, mockValidator)

	t.Run("word exists", func(t *testing.T) {
		assert.True(t, wb.IsExists("word1"))
	})

	t.Run("word does not exist", func(t *testing.T) {
		assert.False(t, wb.IsExists("word4"))
	})

	t.Run("invalid words", func(t *testing.T) {
		mockValidator.On("IsValidWord", invalidWord).Return(false)
		mockValidator.On("IsValidWord", "word").Return(true)

		// Simulate a response with both valid and invalid words
		tsInvalid := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(invalidWord + "\n" + "word"))
		}))
		defer tsInvalid.Close()

		wbInvalid := NewWordsBank(mockRepo, tsInvalid.URL, mockValidator)

		assert.False(t, wbInvalid.IsExists(invalidWord), "Invalid word should not be added")
		assert.True(t, wbInvalid.IsExists("word"), "Valid word should be added")
	})

	t.Run("http request error", func(t *testing.T) {
		tsError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer tsError.Close()

		wbError := NewWordsBank(mockRepo, tsError.URL, mockValidator)

		assert.Empty(t, wbError.repo.GetWords(), "Word bank should be empty on HTTP error")
	})

	t.Run("scanner error", func(t *testing.T) {
		tsScannerError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate a large response that can cause a scanner buffer error
			longString := strings.Repeat("a", 1000000)
			w.Write([]byte(longString))
		}))
		defer tsScannerError.Close()

		wbScannerError := NewWordsBank(mockRepo, tsScannerError.URL, mockValidator)

		assert.Empty(t, wbScannerError.repo.GetWords(), "Word bank should be empty on scanner error")
	})
}
