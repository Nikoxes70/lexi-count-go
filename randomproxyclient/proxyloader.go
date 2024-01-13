package randomproxyclient

import (
	"encoding/json"
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

type ProxyLoader struct {
	apiKey string
}

func NewProxyLoader(apiKey string) *ProxyLoader {
	return &ProxyLoader{apiKey: apiKey}
}

func (pl *ProxyLoader) LoadProxies() ([]string, error) {
	url := "https://proxy.webshare.io/api/v2/proxy/list/?mode=direct&page=1&page_size=100"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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

	var response ProxyResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	proxies := make([]string, 0)
	for _, proxy := range response.Results {
		if proxy.Valid {
			proxies = append(proxies, spew.Sprintf(proxyAddressFormat, proxy.ProxyAddress, proxy.Port))
		}
	}

	return proxies, nil
}
