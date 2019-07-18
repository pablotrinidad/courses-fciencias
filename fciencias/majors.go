package fciencias

import (
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/pablotrinidad/courses-fciencias/models"
)

const pageURL = "licenciatura/Index"

// FetchMajors return a Major slice with the content of the listing webpage
func FetchMajors() []models.Major {
	var majors []models.Major
	document := GetDocument(pageURL)

	digitsRe := regexp.MustCompile(`(\d+)`)

	rawMajors := document.Find("#info-contenido ul li a")
	rawMajors.Each(func(i int, m *goquery.Selection) {

		// Avoid last 2 majors since its information is not complete
		if i == rawMajors.Length()-2 {
			return
		}

		var major models.Major
		major.Name = m.Text()

		// External ID
		href, _ := m.Attr("href")
		major.ExternalID, _ = strconv.Atoi(digitsRe.FindString(href))

		majors = append(majors, major)
	})

	updateMajors(&majors)

	return majors
}

func updateMajors(majors *[]models.Major) {
	db := models.GetDB()
	for _, major := range *majors {
		db.Where(&models.Major{ExternalID: major.ExternalID}).Select("id").First(&major)
		if major.ID == 0 {
			db.Create(&major)
		} else {
			db.Save(&major)
		}
	}
}
