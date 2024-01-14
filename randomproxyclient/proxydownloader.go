package randomproxyclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"net/http"
)

const proxyAddressFormat = "%s:%d"

type Proxy struct {
	ProxyAddress string `json:"proxy_address"`
	Port         int    `json:"port"`
	Valid        bool   `json:"valid"`
}

type ProxyResponse struct {
	Results []Proxy `json:"results"`
}

type ProxyDownloader struct {
	apiKey string
}

func NewProxyDownoader(apiKey string) *ProxyDownloader {
	return &ProxyDownloader{apiKey: apiKey}
}

func (pl *ProxyDownloader) DownloadProxies() ([]string, error) {
	url := "https://proxy.webshare.io/api/v2/proxy/list/?mode=direct&page=1&page_size=100"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+pl.apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseProxies(body)
}

func (pl *ProxyDownloader) RefreshProxies() ([]string, error) {
	client := &http.Client{}

	reqBody := []byte(`{}`) // Replace with actual request body if needed

	req, err := http.NewRequest("POST", "https://proxy.webshare.io/api/v2/proxy/list/refresh/", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set required headers here
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+pl.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return parseProxies(body)
}

func parseProxies(body []byte) ([]string, error) {
	var response ProxyResponse

	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	proxies := make([]string, 0)
	for _, proxy := range response.Results {
		if proxy.Valid {
			proxies = append(proxies, spew.Sprintf(proxyAddressFormat, proxy.ProxyAddress, proxy.Port))
		}
	}
	return proxies, err
}
