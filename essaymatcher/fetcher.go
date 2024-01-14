package essaymatcher

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var animationChars = []string{"|", "/", "-", "\\"}

type matcherer interface {
	ProcessEssay(url string, wordCounts map[string]int) error
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
	mu := &sync.Mutex{}
	errs := make([]error, 0)
	wordCounts := make(map[string]int)

	workCh := make(chan string, len(urls))
	resultCh := make(chan map[string]int, f.concurrencyLimit)
	errCh := make(chan error, f.concurrencyLimit)
	progressCh := make(chan struct{}, len(urls))

	for i := 0; i < f.concurrencyLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range workCh {
				localWordCounts := make(map[string]int)
				if err := f.matcher.ProcessEssay(url, localWordCounts); err != nil {
					errCh <- err
				} else {
					resultCh <- localWordCounts
				}
				progressCh <- struct{}{}
			}
		}()
	}

	go func() {
		for _, url := range urls {
			workCh <- url
		}
		close(workCh)
	}()

	go func() {
		startTime := time.Now()

		for i := 0; i < len(urls); i++ {
			<-progressCh
			mu.Lock()
			percentComplete := (i + 1) * 100 / len(urls)
			elapsedTime := time.Since(startTime).Round(time.Second)
			animationChar := animationChars[i%len(animationChars)]
			fmt.Printf("\rProcessing essays: [%s%s] %d%% Complete %s | Elapsed Time: %s",
				strings.Repeat("=", percentComplete), strings.Repeat(" ", 100-percentComplete), percentComplete, animationChar, elapsedTime)
			mu.Unlock()

		}
	}()

	go func() {
		for {
			select {
			case words, ok := <-resultCh:
				if !ok {
					resultCh = nil
				} else {
					mu.Lock()
					for word, count := range words {
						wordCounts[word] += count
					}
					mu.Unlock()
				}
			case err, ok := <-errCh:
				if !ok {
					errCh = nil
				} else {
					mu.Lock()
					errs = append(errs, err)
					mu.Unlock()
				}
			}

			if resultCh == nil && errCh == nil {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()

	topWords := f.wcp.FindTopWords(wordCounts, f.topNWords)
	var errStrings []string
	for _, e := range errs {
		if e != nil {
			errStrings = append(errStrings, e.Error())
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
	//lines = lines[:1000]
	return lines, nil
}
