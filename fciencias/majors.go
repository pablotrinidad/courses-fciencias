package fciencias

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/PuerkitoBio/goquery"
	"github.com/pablotrinidad/courses-fciencias/storage"
)

const pageURL = "licenciatura/Index"

var digitsRe = regexp.MustCompile(`(\d+)`)

// FetchMajors return a Major slice with the content of the listing webpage
func FetchMajors() []Major {
	var majors []Major
	document := GetDocument(pageURL)

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

	go updateOnDatastore(&majors)

	return majors
}

func updateOnDatastore(majors *[]Major) {
	client := storage.NewDatastoreClient()
	ctx := context.Background()
	for _, major := range *majors {
		q := datastore.NewQuery("Major").Filter("external_id =", major.ExternalID)
		if c, _ := client.Count(ctx, q); c == 0 {
			major.CreatedAt = time.Now()
			client.Put(ctx, datastore.IncompleteKey("Major", nil), &major)
		}
	}

}
