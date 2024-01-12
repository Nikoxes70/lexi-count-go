package essaymatcher

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
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

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			semaphore <- struct{}{}
			if err = f.matcher.ProcessEssay(url, wordCounts, mu); err != nil {
				errs = append(errs, err)
			}
			<-semaphore
		}(url)
	}
	wg.Wait()

	topWords := f.wcp.FindTopWords(wordCounts, f.topNWords)
	response := struct {
		Succeed []wordCountPair `json:"Succeed"`
		Failed  []error         `json:"Failed"`
	}{
		Succeed: topWords,
		Failed:  errs,
	}

	jsonOutput, err := json.Marshal(response)
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
