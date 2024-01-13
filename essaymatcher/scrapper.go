package essaymatcher

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	titleElement     = "h1"
	ogTitle          = "meta[property='og:headline']"
	contentAttr      = "content"
	contentBodyAttr  = ".caas-body-content p"
	maxNumOfAttempts = 10
)

type randomProxyClient interface {
	NewHTTPClientWithRandomProxy() (*http.Client, error)
}

type Scraper struct {
	client randomProxyClient
}

func NewScraper(c randomProxyClient) *Scraper {
	return &Scraper{
		client: c,
	}
}

func (s *Scraper) Scrap(url string, attempt int) (string, error) {
	doc, err := s.htmlDocument(url, attempt, nil)
	if err != nil {
		return "", err
	}

	if doc == nil {
		return "", nil
	}

	title := s.extractTitle(doc)
	article := s.extractArticle(doc)
	article = strings.Replace(article, title, "", -1)
	subtitle := s.extractSubtitle(doc)
	article = strings.TrimSpace(article)
	article = strings.TrimSuffix(article, "\n")
	article = strings.TrimRight(article, "\r\n")

	return fmt.Sprintf("%s %s %s", title, subtitle, article), nil
}

func (s *Scraper) extractTitle(doc *goquery.Document) string {
	return doc.Find(titleElement).Text()
}

func (s *Scraper) extractSubtitle(doc *goquery.Document) string {
	subHeadline, _ := doc.Find(ogTitle).Attr(contentAttr)
	return subHeadline
}

func (s *Scraper) extractArticle(doc *goquery.Document) string {
	var paragraphs []string
	doc.Find(contentBodyAttr).Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	return strings.Join(paragraphs, "\n")
}

func (s *Scraper) htmlDocument(url string, attempt int, lastError error) (*goquery.Document, error) {
	if url == "" {
		return nil, fmt.Errorf("htmlDocument - url must not be empty")
	}

	if attempt >= maxNumOfAttempts {
		return nil, fmt.Errorf("htmlDocument - maximum retry attempts reached with last error: %v", lastError)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("htmlDocument - client.Do - failed to init request")
	}

	client, err := s.client.NewHTTPClientWithRandomProxy()
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return s.htmlDocument(url, attempt+1, err)
	}

	if resp.StatusCode > 400 {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("page not found - status code: %d", resp.StatusCode)
		}
		return s.htmlDocument(url, attempt+1, fmt.Errorf("failed with status code: %d", resp.StatusCode))
	}

	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("htmlDocument - goquery.NewDocumentFromReader - failed to parse html resp for %v with err[%v]", url, err)
	}

	return document, nil
}
