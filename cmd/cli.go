package main

import (
	"bufio"
	"encoding/json"
	"flag"
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
	ProxyApiKey   string `json:"proxy_api_key"`
	WordsBankURL  string `json:"words_bank"`
	EssaysURL     string `json:"essays_url"`
	Threads       int    `json:"threads"`
	TopNWords     int    `json:"top_n_words"`
}

func main() {
	// Define command line flags
	proxyUsername := flag.String("proxy_username", "", "Proxy Username")
	proxyPassword := flag.String("proxy_password", "", "Proxy Password")
	proxyApiKey := flag.String("proxy_api_key", "", "Proxy Api Key")
	wordsBankURL := flag.String("words_bank", "", "URL of the Words Bank")
	essaysURL := flag.String("essays_url", "", "URL of the Essays")
	threads := flag.Int("threads", 0, "Number of threads")
	topNWords := flag.Int("top_n_words", 0, "Number of top words to fetch")
	configPath := flag.String("config", "config/config.json", "Path to the config file")

	flag.Parse()

	// Load config from file
	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("Error opening configuration file: %v\n", err)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding configuration: %v\n", err)
	}

	// Override config with flags if provided
	if *proxyUsername != "" {
		config.ProxyUsername = *proxyUsername
	}
	if *proxyPassword != "" {
		config.ProxyPassword = *proxyPassword
	}
	if *proxyApiKey != "" {
		config.ProxyApiKey = *proxyApiKey
	}
	if *wordsBankURL != "" {
		config.WordsBankURL = *wordsBankURL
	}
	if *essaysURL != "" {
		config.EssaysURL = *essaysURL
	}
	if *threads != 0 {
		config.Threads = *threads
	}
	if *topNWords != 0 {
		config.TopNWords = *topNWords
	}

	// Check for mandatory flags
	if config.ProxyUsername == "" || config.ProxyPassword == "" {
		fmt.Println("Error: proxy_username and proxy_password are required.")
		flag.Usage()
		os.Exit(1)
	}

	// Load proxies
	proxies, err := loadProxies("config/proxies.txt")
	if err != nil {
		fmt.Printf("Error loading proxies: %v\n", err)
		return
	}

	repo := wordsbank.NewRepo()
	validator := wordsbank.NewValidator()
	wb := wordsbank.NewWordsBank(repo, config.WordsBankURL, validator)
	var pl *randomproxyclient.ProxyDownloader
	if config.ProxyApiKey != "" {
		pl = randomproxyclient.NewProxyDownoader(config.ProxyApiKey)
	}

	client := randomproxyclient.NewRandomProxyClient(proxies, config.ProxyUsername, config.ProxyPassword, pl)
	wcp := essaymatcher.NewWordCountPair()

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
