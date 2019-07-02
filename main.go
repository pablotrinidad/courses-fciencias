// Course FCiencias.
// A web crawler made to download and store UNAM's Faculty of Science
// courses schedules.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL       = "http://www.fciencias.unam.mx"
	httpUserAgent = "SchedulesCrawlerBot v1.0 https://github.com/pablotrinidad/courses-fciencias/ | Download course schedules public on the website."
)

var httpClient = &http.Client{
	Timeout: 2 * time.Second,
}

func main() {
	indexDoc := getDocument("docencia/horarios/indice")
	indexDoc.Find("#info-contenido h2").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%d) %s\n", i+1, s.Text())
	})
}

func getDocument(url string) *goquery.Document {
	url = baseURL + "/" + url
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", httpUserAgent)

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
