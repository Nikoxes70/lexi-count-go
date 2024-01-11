package randomproxyclient

import (
	"crypto/tls"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	httpScheme       = "http"
	timeoutInSeconds = 10
)

type RandomProxyClient struct {
	proxies  []string
	username string
	password string
}

func NewRandomProxyClient(proxies []string, username, password string) *RandomProxyClient {
	return &RandomProxyClient{
		proxies:  proxies,
		username: username,
		password: password,
	}
}

func (rpc *RandomProxyClient) NewHTTPClientWithRandomProxy() (*http.Client, error) {
	customTransport := &http.Transport{
		Proxy: http.ProxyURL(&url.URL{
			Scheme: httpScheme,
			User:   url.UserPassword(rpc.username, rpc.password),
			Host:   rpc.proxies[rand.Intn(len(rpc.proxies))],
		}),
		DialTLS:             httpsDial,
		TLSHandshakeTimeout: timeoutInSeconds * time.Second,
	}

	httpClient := &http.Client{
		Transport: customTransport,
	}

	return httpClient, nil
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
