package randomproxyclient

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	httpScheme       = "http"
	timeoutInSeconds = 10
)

type proxyDownloader interface {
	DownloadProxies() ([]string, error)
	RefreshProxies() ([]string, error)
}

type RandomProxyClient struct {
	proxies        []string
	blockedProxies map[string]string
	username       string
	password       string
	proxyLoader    proxyDownloader
}

func NewRandomProxyClient(proxies []string, username, password string, proxyLoader proxyDownloader) *RandomProxyClient {
	if proxyLoader != nil {
		if newProxies, err := proxyLoader.DownloadProxies(); err == nil {
			proxies = newProxies
		}
	}
	return &RandomProxyClient{
		proxies:        proxies,
		blockedProxies: map[string]string{},
		username:       username,
		password:       password,
		proxyLoader:    proxyLoader,
	}
}

func (rpc *RandomProxyClient) NewHTTPClientWithRandomProxy() (*http.Client, string, error) {
	proxy, err := rpc.randomProxy(1)
	if err != nil {
		newProxies, err := rpc.proxyLoader.RefreshProxies()
		if err != nil {
			return nil, "", err
		}
		rpc.proxies = newProxies
	}

	customTransport := &http.Transport{
		Proxy: http.ProxyURL(&url.URL{
			Scheme: httpScheme,
			User:   url.UserPassword(rpc.username, rpc.password),
			Host:   proxy,
		}),
		DialTLS:             httpsDial,
		TLSHandshakeTimeout: timeoutInSeconds * time.Second,
	}

	httpClient := &http.Client{
		Transport: customTransport,
	}

	return httpClient, proxy, nil
}

func (rpc *RandomProxyClient) MarkProxyAsBlocked(proxyURL string) {
	mu := sync.Mutex{}
	mu.Lock()
	rpc.blockedProxies[proxyURL] = proxyURL
	mu.Unlock()
}

func (rpc *RandomProxyClient) isProxyBlocked(proxyURL string) bool {
	_, exists := rpc.blockedProxies[proxyURL]
	return exists
}

func (rpc *RandomProxyClient) randomProxy(attempt int) (string, error) {
	if attempt >= len(rpc.proxies) {
		return "", fmt.Errorf("all proxies are blocked or attempts exceeded")
	}

	p := rpc.proxies[rand.Intn(len(rpc.proxies))]
	if _, ok := rpc.blockedProxies[p]; ok {
		return rpc.randomProxy(attempt + 1)
	}
	return p, nil
}

func httpsDial(network, addr string) (net.Conn, error) {
	targetConn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	tlsConn := tls.Client(targetConn, &tls.Config{})

	if err = tlsConn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return nil, err
	}

	if err := tlsConn.Handshake(); err != nil {
		tlsConn.Close()
		return nil, err
	}

	if err = tlsConn.SetDeadline(time.Time{}); err != nil {
		return nil, err
	}

	return tlsConn, nil
}
