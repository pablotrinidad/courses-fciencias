package crawler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// callConcurrent perform fns concurrently.
func callConcurrent(fns []func()) {
	var wg sync.WaitGroup
	for i := range fns {
		f := fns[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
}

// getDocument fetches the given URL and return the corresponding goquery Document
func getDocument(url string) (*goquery.Document, error) {
	httpClient := &http.Client{Timeout: requestTimeout}
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
