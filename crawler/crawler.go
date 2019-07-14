// Package crawler contains utility functions for fetching and parsing web pages
package crawler

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// BaseURL is the main website path from where all pages will be downloaded
	BaseURL = "http://www.fciencias.unam.mx"

	// HTTPUserAgent is the string sent in the User-Agent header on every request
	HTTPUserAgent = "SchedulesCrawlerBot v1.0 https://github.com/pablotrinidad/courses-fciencias/ | Download course schedules public on the website."
)

var httpClient = &http.Client{
	Timeout: 2 * time.Second,
}

// GetDocument fetches the given path and return a goquery Document with its content
func GetDocument(url string) *goquery.Document {
	url = BaseURL + "/" + url
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", HTTPUserAgent)

	// Perform request
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Handle unsuccessful request
	if response.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load document using goquery
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return document
}
