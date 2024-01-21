# LexiCount CLI Tool

## Overview
This application, written in Go (Golang), is designed to fetch a list of essays from a provided source and analyze these essays to determine the top 10 most frequently occurring words. The application adheres to specific word validation rules and outputs the results in a well-structured, pretty-printed JSON format.

### Features
`Concurrency`: Utilizes Go's concurrency features for efficient processing.

`Efficiency`: Optimized for quick retrieval and processing of data.

`Rate Limiting`: Ensures controlled access to resources, preventing overloading.

`Comprehensibility`: Code is clear, well-documented, and adheres to Go's coding standards.

### Word Validation Rules
`A word is considered valid if it:`

Contains at least 3 characters.

Contains only alphabetic characters.

Is included in a predefined bank of words.

## Getting Started

## Dependencies
### [Webshare.io](webshare.io) 
LexiCount requires an active and registered user account from Webshare.io to function correctly. The tool utilizes multiple proxies provided by Webshare.io to scrape a vast list of essay URLs concurrently. This approach helps in efficient data retrieval and minimizes the chances of being rate-limited or blocked by target websites.

Ensure that you have a registered account with Webshare.io and access to a list of proxy servers. Your proxy credentials (username and password) and the proxy list should be included in the tool's configuration as per the instructions below.



## Installation

Before installing LexiCount, ensure you have Go installed on your system. If not, [download and install Go](https://golang.org/dl/).

To install LexiCount, clone the repository and build the binary:

```bash
git clone https://github.com/Nikoxes70/lexi-count-go
cd lexi-count-go 
go build cmd/cli.go
```

Run LexiCount with the required flags for proxy_username and proxy_password for [WEBSHARE](webshare.io).
```bash
./cli --proxy_username="username" --proxy_password="password"
```

Run LexiCount with the prefered flags for proxy_username, proxy_password and apiKey for [WEBSHARE](webshare.io).
```bash
./cli --proxy_username="username" --proxy_password="password" --proxy_api_key="apiKey"
```


Optional flags can override other settings:
```bash
./cli --proxy_username="username" --proxy_password="password" --proxy_api_key="apiKey" --threads=10 --top_n_words=20
```

#### Flags

| Name           | Type   | Description                  |
|----------------|--------|------------------------------|
| proxy_username | string | webshare proxy Username      |
| proxy_password | string | webshare proxy Password      |
| proxy_api_key  | string | webshare Proxy Api Key       |
| words_bank     | string | URL of the Words Bank        |
| essays_url     | stirng | URL of the Essays            |
| threads        | int    | Number of threads            |
| top_n_words    | int    | Number of top words to fetch |