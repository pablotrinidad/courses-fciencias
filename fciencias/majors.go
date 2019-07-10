// Package fciencias contains the functions used to parse the majors UNAM's
// Faculty of Science run.
package fciencias

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pablotrinidad/courses-fciencias/crawler"
)

// GetMajors return all majors listed in the website.
func GetMajors() {
	// Match content inside parenthesis
	re := regexp.MustCompile(`\([^)]+\)`)

	// Fetch main index document
	indexDoc := crawler.GetDocument("docencia/horarios/indice")

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
