package fciencias

import (
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const pageURL = "licenciatura/Index"

// FetchMajors return a Major slice with the content of the listing webpage
func FetchMajors() []Major {
	var majors []Major
	document := GetDocument(pageURL)

	digitsRe := regexp.MustCompile(`(\d+)`)

	rawMajors := document.Find("#info-contenido ul li a")
	rawMajors.Each(func(i int, m *goquery.Selection) {

		// Avoid last 2 majors since its information is not complete
		if i == rawMajors.Length()-2 {
			return
		}

		var major Major
		major.Name = m.Text()

		// External ID
		href, _ := m.Attr("href")
		major.ExternalID, _ = strconv.Atoi(digitsRe.FindString(href))

		majors = append(majors, major)
	})

	return majors
}
