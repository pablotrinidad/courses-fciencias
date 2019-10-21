package pages

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/pablotrinidad/courses-fciencias/crawler"
	"github.com/pablotrinidad/courses-fciencias/crawler/entities"
)

func FetchProgramCourses(program *entities.Program) (data []*entities.ProgramCourse) {
	document := crawler.GetDocument(program.GetURL())
	log.Println(program.GetURL())

	parenthesisRegex := regexp.MustCompile(`\(([^\)]+)\)`)
	digitRegex := regexp.MustCompile(`\d+`)

	programRawName := document.Find("h1").First().Text()
	programRawName = parenthesisRegex.FindString(programRawName)
	program.Name = strings.Title(programRawName[1 : len(programRawName)-1])

	program.Year, _ = strconv.Atoi(digitRegex.FindString(program.Name))

	uls := document.Find("#info-contenido ul").Last().Find("p,h3,h2")
	semester, mandatory := 0, true
	courseNameRegex := regexp.MustCompile(`, \d+ créditos\.`)

	uls.Each(func(i int, s *goquery.Selection) {
		switch {
		case s.Is("h3"):
			semester++
			fmt.Printf("%d (%s)\n", semester, s.Text())
		case s.Is("h2") && i > 0:
			mandatory = false
			semester = -1
			fmt.Println("OPTATIVAS")
		case s.Is("p") && s.Find("a").Length() != 0:
			rawText := strings.TrimSpace(s.Text())
			creditsLocation, cutIndex := courseNameRegex.FindStringIndex(rawText), len(rawText)

			credits := 0
			if len(creditsLocation) > 0 {
				cutIndex = creditsLocation[0]
				credits, _ = strconv.Atoi(digitRegex.FindString(rawText[cutIndex:]))
			}

			course := entities.ProgramCourse{
				BaseEntity: entities.BaseEntity{ExternalID: 0},
				Program:    program.ExternalID,
				Name:       rawText[:cutIndex],
				Semester:   semester,
				Credits:    credits,
				Syllabus:   "",
				Mandatory:  mandatory,
			}

			courseURL, ok := s.Find("a").First().Attr("href")
			if ok {
				seps := strings.Split(courseURL, "/")
				course.ExternalID, _ = strconv.Atoi(seps[len(seps)-1])
			}

			fmt.Printf("\t%d) %s - %d credits\n", course.ExternalID, course.Name, course.Credits)
		}
	})

	return data
}

// FetchMajorCourses
func FetchMajorCourses(major int) (map[int][]*entities.ProgramCourse, map[int]*entities.Program) {
	courses := make(map[int][]*entities.ProgramCourse)
	programs := make(map[int]*entities.Program)

	var wg sync.WaitGroup
	cn := make(chan *entities.ProgramCourse)

	for _, programID := range entities.Programs[major] {
		wg.Add(1)
		programs[programID] = &entities.Program{
			BaseEntity: entities.BaseEntity{ExternalID: programID},
			Major:      major,
			Name:       "",
			Year:       0,
		}
		go func(id int, cn chan *entities.ProgramCourse, program *entities.Program) {
			defer wg.Done()
			for _, course := range FetchProgramCourses(program) {
				cn <- course
			}
		}(programID, cn, programs[programID])
		break
	}

	go func() {
		wg.Wait()
		close(cn)
	}()
	for course := range cn {
		courses[course.Program] = append(courses[course.Program], course)
	}

	return courses, programs
}
