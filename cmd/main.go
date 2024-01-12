package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"LexiCount/essaymatcher"
	"LexiCount/randomproxyclient"
	"LexiCount/wordsbank"
)

type Config struct {
	ProxyUsername string `json:"proxy_username"`
	ProxyPassword string `json:"proxy_password"`
	WordsBankURL  string `json:"words_bank"`
	EssaysURL     string `json:"essays_url"`
	Threads       int    `json:"threads"`
	TopNWords     int    `json:"top_n_words"`
}

func main() {
	configFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Printf("Error opening configuration file: %v\n", err)
		os.Exit(1)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("Error decoding configuration: %v\n", err)
		os.Exit(1)
	}

	proxies, err := loadProxies("config/proxies.txt")
	if err != nil {
		fmt.Printf("Error loading proxies: %v\n", err)
		return
	}

	repo := wordsbank.NewRepo()
	validator := wordsbank.NewValidator()
	wb := wordsbank.NewWordsBank(repo, config.WordsBankURL, validator)

	wcp := essaymatcher.NewWordCountPair()
	client := randomproxyclient.NewRandomProxyClient(proxies, config.ProxyUsername, config.ProxyPassword)

	s := essaymatcher.NewScraper(client)
	em := essaymatcher.NewEssayMatcher(validator, s, wb)

	fetcher := essaymatcher.NewFetcher(config.Threads, config.TopNWords, em, wcp)
	r, err := fetcher.Start(config.EssaysURL)
	if err != nil {
		log.Fatal(err)
	}
	println(r)
}

func loadProxies(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var proxies []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}
	return proxies, scanner.Err()
}
