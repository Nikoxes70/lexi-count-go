package essaymatcher

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMatcherer mocks the matcherer interface
type MockMatcherer struct {
	mock.Mock
}

func (m *MockMatcherer) ProcessEssay(url string, wordCounts map[string]int, mu *sync.Mutex) error {
	args := m.Called(url, wordCounts, mu)
	return args.Error(0)
}

// MockWordCountPairer mocks the wordCountPairer interface
type MockWordCountPairer struct {
	mock.Mock
}

func (m *MockWordCountPairer) FindTopWords(wordCounts map[string]int, n int) []wordCountPair {
	args := m.Called(wordCounts, n)
	return args.Get(0).([]wordCountPair)
}

// TestFetcher_Start tests the Start function of Fetcher
func TestFetcher_Start(t *testing.T) {
	mockMatcher := new(MockMatcherer)
	mockWCP := new(MockWordCountPairer)

	// Set up a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("https://example.com/essay1\nhttps://example.com/essay2"))
	}))
	defer ts.Close()

	fetcher := NewFetcher(2, 5, mockMatcher, mockWCP)

	t.Run("successful fetch", func(t *testing.T) {
		mockMatcher.On("ProcessEssay", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		mockWCP.On("FindTopWords", mock.Anything, 5).Return([]wordCountPair{{Word: "test", Count: 10}})

		result, err := fetcher.Start(ts.URL)
		assert.NoError(t, err)
		assert.Contains(t, result, "\"Succeed\":[{\"Word\":\"test\",\"Count\":10}]")
	})

	t.Run("http request error", func(t *testing.T) {
		tsError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer tsError.Close()

		_, err := fetcher.Start(tsError.URL)
		assert.Error(t, err)
	})

	t.Run("concurrency handling", func(t *testing.T) {
		// Modify the mock matcher to simulate concurrent processing
		mockMatcher.On("ProcessEssay", mock.Anything, mock.Anything, mock.Anything).Return(func(url string, wordCounts map[string]int, mu *sync.Mutex) error {
			mu.Lock()
			defer mu.Unlock()
			wordCounts["test"]++
			return nil
		})

		mockWCP.On("FindTopWords", mock.Anything, 5).Return([]wordCountPair{{Word: "test", Count: 2}})

		result, err := fetcher.Start(ts.URL)
		assert.NoError(t, err)
		assert.Contains(t, result, "\"Succeed\":[{\"Word\":\"test\",\"Count\":10}]")
	})

	t.Run("no urls fetched", func(t *testing.T) {
		tsEmpty := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(""))
		}))
		defer tsEmpty.Close()

		result, err := fetcher.Start(tsEmpty.URL)
		assert.NoError(t, err)
		assert.Contains(t, result, "{\"Succeed\":[{\"Word\":\"test\",\"Count\":10}]")
		assert.Contains(t, result, "\"Failed\":[]")
	})
}
