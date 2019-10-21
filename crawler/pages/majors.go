package pages

import (
	"github.com/pablotrinidad/courses-fciencias/crawler"
	"github.com/pablotrinidad/courses-fciencias/crawler/entities"
	"strings"
	"sync"
)

// FetchMajor download a major's website and parse the input
func FetchMajor(id int) entities.Major {
	major := entities.Major{}
	major.ExternalID = id

	document := crawler.GetDocument(major.GetURL())

	rawName := document.Find("h1").First().Text()
	major.Name = strings.Title(strings.TrimSpace(strings.Split(rawName, "(")[0]))

	return major
}

// FetchAllMajors concurrently. It uses the FetchMajor function.
func FetchAllMajors() []*entities.Major {
	var wg sync.WaitGroup
	cn := make(chan *entities.Major, len(entities.Majors))

	for _, major := range entities.Majors {
		wg.Add(1)
		go func(id int, cn chan *entities.Major) {
			defer wg.Done()
			major := FetchMajor(id)
			cn <- &major
		}(major, cn)
	}

	wg.Wait()
	close(cn)

	majors, i := make([]*entities.Major, len(entities.Majors)), 0
	for major := range cn {
		majors[i] = major
		i++
	}

	return majors
}
