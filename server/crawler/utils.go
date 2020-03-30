package crawler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var httpClient = &http.Client{Timeout: 2 * time.Second}

// callConcurrent perform fns concurrently.
func callConcurrent(fns ...func()) {
	var wg sync.WaitGroup
	for _, fn := range fns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}

// getDocument fetches the given URL and return the corresponding goquery Document
func getDocument(url string) (*goquery.Document, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new HTTP request; %v", err)
	}
	request.Header.Set("User-Agent", httpUserAgent)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request; %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("got failed success code %d", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse page; %v", err)
	}
	return document, nil
}
