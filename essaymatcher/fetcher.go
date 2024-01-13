package essaymatcher

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type matcherer interface {
	ProcessEssay(url string, wordCounts map[string]int, mu *sync.Mutex) error
}

type wordCountPairer interface {
	FindTopWords(wordCounts map[string]int, n int) []wordCountPair
}

type Fetcher struct {
	concurrencyLimit int
	topNWords        int
	matcher          matcherer
	wcp              wordCountPairer
}

func NewFetcher(concurrencyLimit int, topNWords int, matcher matcherer, wcp wordCountPairer) *Fetcher {
	return &Fetcher{concurrencyLimit: concurrencyLimit, topNWords: topNWords, matcher: matcher, wcp: wcp}
}

func (f *Fetcher) Start(essaysUrl string) (string, error) {
	urls, err := fetchURLs(essaysUrl)
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	wordCounts := make(map[string]int)
	errs := make([]error, 0)
	mu := &sync.Mutex{}
	semaphore := make(chan struct{}, f.concurrencyLimit)

	totalURLs := len(urls)
	var completed int

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Temporary variable for storing word counts
			localWordCounts := make(map[string]int)

			// Process the essay and store results in localWordCounts
			if err = f.matcher.ProcessEssay(url, localWordCounts, nil); err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}

			// Safely update the shared wordCounts map
			mu.Lock()
			for word, count := range localWordCounts {
				wordCounts[word] += count
			}

			completed++
			percentComplete := (completed * 100) / totalURLs
			fmt.Printf("\rProcessing essays: [%s%s] %d%% Complete",
				strings.Repeat("=", percentComplete), strings.Repeat(" ", 100-percentComplete), percentComplete)

			mu.Unlock()
		}(url)
	}
	wg.Wait()

	topWords := f.wcp.FindTopWords(wordCounts, f.topNWords)
	var errStrings []string
	for _, err := range errs {
		if err != nil {
			errStrings = append(errStrings, err.Error())
		}
	}
	response := struct {
		Succeed []wordCountPair `json:"Succeed"`
		Failed  []string        `json:"Failed"`
	}{
		Succeed: topWords,
		Failed:  errStrings,
	}

	jsonOutput, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsonOutput), nil
}

func fetchURLs(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-ok HTTP status: %s", resp.Status)
	}

	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return lines, nil
}
