# LexiCount CLI Tool

LexiCount is a command-line tool designed for processing and analyzing essays. It leverages proxies for network requests and analyzes text content to identify top words used in a collection of essays.

## Dependencies
### Webshare.io 
LexiCount requires an active and registered user account from Webshare.io to function correctly. The tool utilizes multiple proxies provided by Webshare.io to scrape a vast list of essay URLs concurrently. This approach helps in efficient data retrieval and minimizes the chances of being rate-limited or blocked by target websites.

Ensure that you have a registered account with Webshare.io and access to a list of proxy servers. Your proxy credentials (username and password) and the proxy list should be included in the tool's configuration as per the instructions below.



## Installation

Before installing LexiCount, ensure you have Go installed on your system. If not, [download and install Go](https://golang.org/dl/).

To install LexiCount, clone the repository and build the binary:

```bash
git clone https://github.com/your-repository/LexiCount.git
cd LexiCount
go build
```

Run LexiCount with the required flags for proxy_username and proxy_password for [WEBSHARE](webshare.io).
```bash
./LexiCount --proxy_username="username" --proxy_password="password"
```

Optional flags can override other settings:
```bash
./LexiCount --proxy_username="username" --proxy_password="password" --threads=10 --top_n_words=20
```




