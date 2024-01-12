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
	mu := &sync.Mutex{}

	semaphore := make(chan struct{}, f.concurrencyLimit)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			semaphore <- struct{}{}
			err = f.matcher.ProcessEssay(url, wordCounts, mu)
			if err != nil {
				fmt.Printf("🔴Error processEssay: %v \n", err)
			}
			<-semaphore
		}(url)
	}
	wg.Wait()

	topWords := f.wcp.FindTopWords(wordCounts, f.topNWords)
	jsonOutput, err := json.MarshalIndent(topWords, "", "  ")
	if err != nil {
		panic(err)
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
