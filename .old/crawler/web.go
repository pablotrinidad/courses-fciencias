package crawler

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	BaseURL       = "http://www.fciencias.unam.mx"
	HTTPUserAgent = "CoursesCrawlerBot v.1.0 https://github.com/pablotrinidad/courses-fciencias | Download courses catalog"
)

var httpClient = &http.Client{Timeout: 2 * time.Second}

// GetDocument fetches the given URL and return the corresponding goquery Document
func GetDocument(url string) *goquery.Document {
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
