// Course FCiencias.
// A web crawler made to download and store UNAM's Faculty of Science
// courses schedules.

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
	// Match content inside parenthesis
	re := regexp.MustCompile(`\([^)]+\)`)

	// Fetch main index document
	indexDoc := getDocument("docencia/horarios/indice")

	// Obtain careers and their academic plans
	indexDoc.Find("#info-contenido h2").Each(func(i int, s *goquery.Selection) {
		// Degree name
		degreeName := s.Text()
		fmt.Printf("%d) %s\n", i+1, degreeName)

		// Academic plans
		s.Next().Find("a").Each(func(j int, t *goquery.Selection) {
			academicPlan := re.FindString(t.Text())
			academicPlan = strings.Title(academicPlan[1 : len(academicPlan)-1])

			// Obtain plan ID from URL
			ref, _ := t.Attr("href")
			planURL, err := url.Parse(ref)
			if err != nil {
				log.Fatal(err)
			}
			urlComponents := strings.Split(planURL.Path, "/")
			academicPlanID, err := strconv.Atoi(urlComponents[len(urlComponents)-1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\t%d.%d) % s (with ID: %d)\n", i, j, academicPlan, academicPlanID)
		})
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
