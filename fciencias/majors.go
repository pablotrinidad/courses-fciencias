package fciencias

import (
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const pageURL = "licenciatura/Index"

var digitsRe = regexp.MustCompile(`(\d+)`)

// FetchMajors return a Major slice with the content of the listing webpage
func FetchMajors() []Major {
	var majors []Major
	document := GetDocument(pageURL)

	rawMajors := document.Find("#info-contenido ul li a")
	rawMajors.Each(func(i int, m *goquery.Selection) {

		// Avoid last major since its information is not complete
		if i == rawMajors.Length()-1 {
			return
		}

		var major Major
		major.ID = i + 1
		major.Name = m.Text()

		// External ID
		href, _ := m.Attr("href")
		major.ExternalID, _ = strconv.Atoi(digitsRe.FindString(href))

		majors = append(majors, major)
	})

	return majors
}
