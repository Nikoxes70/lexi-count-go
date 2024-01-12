package essaymatcher

import (
	"fmt"
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
		assert.Contains(t, result, "test")
	})

	t.Run("http request failure", func(t *testing.T) {
		// Set up a mock HTTP server that returns an error
		tsErr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer tsErr.Close()

		result, err := fetcher.Start(tsErr.URL)
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("matcherer processing error", func(t *testing.T) {
		mockMatcher.On("ProcessEssay", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("processing error"))
		mockWCP.On("FindTopWords", mock.Anything, 5).Return([]wordCountPair{})

		_, err := fetcher.Start(ts.URL)
		assert.Error(t, err)
		// You can also assert that the error message contains "processing error"
	})

	t.Run("invalid url handling", func(t *testing.T) {
		// Use a mock server that returns an invalid URL
		tsInvalidURL := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("http://%zzz"))
		}))
		defer tsInvalidURL.Close()

		result, err := fetcher.Start(tsInvalidURL.URL)
		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("invalid url handling", func(t *testing.T) {
		// Use a mock server that returns an invalid URL
		tsInvalidURL := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("http://%zzz"))
		}))
		defer tsInvalidURL.Close()

		result, err := fetcher.Start(tsInvalidURL.URL)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
