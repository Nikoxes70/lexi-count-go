package essaymatcher

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockValidatorer mocks the validatorer interface
type MockValidatorer struct {
	mock.Mock
}

func (m *MockValidatorer) IsValidWord(word string) bool {
	args := m.Called(word)
	return args.Bool(0)
}

// MockScrapperer mocks the scrapperer interface
type MockScrapperer struct {
	mock.Mock
}

func (m *MockScrapperer) Scrap(url string, attempt int) (string, error) {
	args := m.Called(url, attempt)
	return args.String(0), args.Error(1)
}

// MockWordsBank mocks the wordsBank interface
type MockWordsBank struct {
	mock.Mock
}

func (m *MockWordsBank) IsExists(word string) bool {
	args := m.Called(word)
	return args.Bool(0)
}

// TestEssayMatcher_ProcessEssay tests the ProcessEssay function of EssayMatcher
func TestEssayMatcher_ProcessEssay(t *testing.T) {
	mockValidator := new(MockValidatorer)
	mockScrapper := new(MockScrapperer)
	mockWordsBank := new(MockWordsBank)

	essayMatcher := NewEssayMatcher(mockValidator, mockScrapper, mockWordsBank)

	// Mock the dependencies
	mockScrapper.On("Scrap", "http://example.com/essay", 1).Return("This is a test essay", nil)
	mockValidator.On("IsValidWord", mock.Anything).Return(true)
	mockWordsBank.On("IsExists", mock.Anything).Return(true)

	wordCounts := make(map[string]int)
	mu := &sync.Mutex{}

	err := essayMatcher.ProcessEssay("http://example.com/essay", wordCounts, mu)

	assert.NoError(t, err)
	assert.NotEmpty(t, wordCounts)

	t.Run("scrapper error", func(t *testing.T) {
		mockScrapper.On("Scrap", "http://example.com/failed-essay", 1).Return("", fmt.Errorf("scrap error"))
		err = essayMatcher.ProcessEssay("http://example.com/failed-essay", wordCounts, mu)
		assert.Error(t, err)
		assert.Equal(t, "scrap error", err.Error())
	})

	t.Run("invalid words", func(t *testing.T) {
		mockScrapper.On("Scrap", "http://example.com/invalid-essay", 1).Return("a bb ccc", nil)
		mockValidator.On("IsValidWord", "a").Return(false)
		mockValidator.On("IsValidWord", "bb").Return(false)
		mockValidator.On("IsValidWord", "ccc").Return(true)
		mockWordsBank.On("IsExists", "ccc").Return(true)

		wordCounts = make(map[string]int)
		err = essayMatcher.ProcessEssay("http://example.com/invalid-essay", wordCounts, mu)

		assert.NoError(t, err)
		assert.Equal(t, 1, wordCounts["ccc"])
		assert.Len(t, wordCounts, 1)
	})

	t.Run("words not in words bank", func(t *testing.T) {
		uniqueword := "uniqueword"
		mockScrapper.On("Scrap", "http://example.com/unique-essay", 1).Return("", nil)
		mockValidator.On("IsValidWord", uniqueword).Return(true)
		mockWordsBank.On("IsExists", uniqueword).Return(false)

		wordCounts = make(map[string]int)
		err = essayMatcher.ProcessEssay("http://example.com/unique-essay", wordCounts, mu)

		assert.NoError(t, err)
		assert.Empty(t, wordCounts)
	})

	t.Run("concurrent processing", func(t *testing.T) {
		mockScrapper.On("Scrap", mock.Anything, 1).Return("word", nil)
		mockValidator.On("IsValidWord", "word").Return(true)
		mockWordsBank.On("IsExists", "word").Return(true)

		wordCounts = make(map[string]int)
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_ = essayMatcher.ProcessEssay("http://example.com/concurrent-essay", wordCounts, mu)
			}()
		}
		wg.Wait()

		assert.Equal(t, 10, wordCounts["word"])
	})
}
